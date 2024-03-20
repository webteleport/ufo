package term

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
	"k0s.io/pkg/wrap"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func wsfy(h string) string {
	switch h {
	case "http":
		return "ws"
	case "https":
		return "wss"
	default:
		return ""
	}
}

func wettyHandler() http.Handler {
	var err error
	shell := os.Getenv("SHELL")

	if shell == "" {
		switch runtime.GOOS {
		case "windows":
			shell = "powershell.exe"
			shell, err = exec.LookPath(shell)
			if err != nil {
				shell, _ = exec.LookPath("cmd.exe")
			}
		default:
			shell = "bash"
			_, err = exec.LookPath(shell)
			if err != nil {
				shell = "sh"
			}
		}
	}
	cmd := []string{shell}

	return &auto{
		fac: factory.New(cmd),
		rp:  utils.ReverseProxy("https://wetty.deno.dev"),
	}
}

func Run(args []string) error {
	handler := wettyHandler()
	arg0 := Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler)
}

type auto struct {
	fac agent.TtyFactory
	nth int
	rp  http.Handler
}

func (a *auto) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isWS := strings.Split(r.Header.Get("Upgrade"), ",")[0] == "websocket"
	if !isWS {
		a.rp.ServeHTTP(w, r)
		return
	}
	a.nth += 1
	conn, err := wrap.Wrconn(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	a.serveConn(conn, a.nth)
}

func (a *auto) serveConn(conn net.Conn, nth int) {
	var (
		tryCommandOnce = &sync.Once{}
		cmdCh          = make(chan []string, 1)
		envCh          = make(chan map[string]string, 1)
		resizeCh       = make(chan struct{ rows, cols int }, 4)
	)

	server := asciitransport.Server(conn)
	// send
	// case output:

	// recv
	go func() {
		for {
			var (
				re   = <-server.ResizeEvent()
				rows = int(re.Height)
				cols = int(re.Width)
			)
			tryCommandOnce.Do(func() {
				cmdCh <- re.Command
				envCh <- re.Env
			})
			resizeCh <- struct{ rows, cols int }{rows, cols}
		}
		server.Close()
	}()

	cmd := <-cmdCh
	env := <-envCh

	_ = cmd
	_ = env

	term, err := a.fac.MakeTty()
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			re := <-resizeCh
			err := term.Resize(re.rows, re.cols)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	logname := fmt.Sprintf("/tmp/term-%d.log", nth)
	logfile, err := os.Create(logname)
	if err == nil {
		defer func() {
			exec.Command("dkg-push", logname).Run()
			log.Println("log written to", logname)
		}()
	}

	opts := []asciitransport.Opt{
		asciitransport.WithReader(term),
		asciitransport.WithWriter(term),
		asciitransport.WithLogger(logfile),
	}
	server.ApplyOpts(opts...)

	<-server.Done()
	term.Close()
}
