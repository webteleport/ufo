//go:build !js && !wasip1

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
		tm:  map[string]agent.Tty{},
	}
}

type auto struct {
	fac agent.TtyFactory
	nth int
	rp  http.Handler
	tm  map[string]agent.Tty
}

func (a *auto) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isWS := strings.Split(r.Header.Get("Upgrade"), ",")[0] == "websocket"
	isUI := os.Getenv("UI") != ""
	if !isWS {
		if isUI {
			a.rp.ServeHTTP(w, r)
		} else {
			http.Error(w, "OK", 200)
		}
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

func (a *auto) makeTerm(cmd []string, env map[string]string, ses string) (agent.Tty, error) {
	if ses == "" {
		term, err := a.fac.MakeTty()
		if err != nil {
			return nil, err
		}
		return term, nil
	} else {
		if term, ok := a.tm[ses]; ok {
			log.Println("old term", ses)
			return term, nil
		}
		term, err := a.fac.MakeTty()
		if err != nil {
			return nil, err
		}
		log.Println("new term", ses, len(a.tm))
		a.tm[ses] = term
		return term, nil
	}
}

func (a *auto) serveConn(conn net.Conn, nth int) {
	var (
		tryCommandOnce = &sync.Once{}
		cmdCh          = make(chan []string, 1)
		envCh          = make(chan map[string]string, 1)
		sesCh          = make(chan string, 1)
		resizeCh       = make(chan struct{ rows, cols int }, 4)
	)

	server := asciitransport.Server(conn)
	// send
	// case output:

	// recv
	go func() {
	Loop:
		for {
			select {
			case re := <-server.ResizeEvent():
				var (
					rows = int(re.Height)
					cols = int(re.Width)
				)
				tryCommandOnce.Do(func() {
					cmdCh <- re.Command
					envCh <- re.Env
					sesCh <- re.Title
				})
				resizeCh <- struct{ rows, cols int }{rows, cols}
			case <-server.Done():
				break Loop
			}
		}
		server.Close()
	}()

	cmd := <-cmdCh
	env := <-envCh
	ses := <-sesCh

	term, err := a.makeTerm(cmd, env, ses)
	if err != nil {
		log.Println(err)
		return
	}

	if ses == "" {
		defer term.Close()
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
}
