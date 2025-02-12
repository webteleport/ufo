package handler

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/btwiuse/wsconn"
	"github.com/webteleport/utils"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
)

func Handler() http.Handler {
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
	conn, err := wsconn.Wrconn(w, r)
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

	opts := []asciitransport.Opt{
		asciitransport.WithReader(term),
		asciitransport.WithWriter(term),
	}
	server.ApplyOpts(opts...)

	<-server.Done()
	term.Close()
}
