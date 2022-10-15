package term

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/webteleport/webteleport"
	"k0s.io"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
	"k0s.io/pkg/reverseproxy"
	"k0s.io/pkg/wrap"
	"nhooyr.io/websocket"
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

func Run(args []string) error {
	ln, err := webteleport.Listen(context.Background(), Arg0(args, "https://ufo.k0s.io"))
	if err != nil {
		return err
	}
	// log.Println("🛸 listening on", wsfy(ln.Network())+"://"+ln.String())
	log.Println("🛸 listening on", ln.ClickableURL())

	return http.Serve(ln, &auto{
		fac: factory.New([]string{"bash"}),
		rp:  reverseproxy.Handler("https://wetty.vercel.app"),
	})
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
	conn, err := wrconn(w, r)
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

	if len(cmd) == 0 {
		cmd = []string{"bash"}
	}

	term, err := a.fac.MakeTtyEnv(cmd, env)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		re := <-resizeCh
		err := term.Resize(re.rows, re.cols)
		if err != nil {
			log.Println(err)
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

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
		// here without the Subprotocols info, Chrome/Edge won't work
		// so must add this piece of info here
		Subprotocols: []string{"wetty"},
	})
	if err != nil {
		return nil, err
	}
	wsconn.SetReadLimit(k0s.MAX_WS_MESSAGE)
	conn := wrap.NetConn(wsconn)
	addr := wrap.NewAddr("websocket", r.RemoteAddr)
	conn = wrap.ConnWithAddr(conn, addr)
	return conn, nil
}