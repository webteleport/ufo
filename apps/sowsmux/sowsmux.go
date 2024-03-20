// $ websocat tcp-listen:0.0.0.0:3000 --binary wss://35.ufo.k0s.io -E
// $ export HTTP_PROXY=http://127.0.0.1:3000 HTTPS_PROXY=http://127.0.0.1:3000
// $ curl -v http://google.com
// $ curl -v https://google.com
package sowsmux

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/btwiuse/gost"
	"github.com/hashicorp/yamux"
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport"

	"k0s.io/pkg/wrap"
)

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
	ln, err := webteleport.Listen(context.Background(), apps.Arg0(args, apps.RELAY))
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", wsfy(ln.Addr().Network())+"://"+ln.Addr().String())

	return http.Serve(ln, utils.GinLoggerMiddleware(&auto{}))
}

type auto struct{}

func (a *auto) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wrap.Wrconn(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	// Setup server side of yamux
	session, err := yamux.Server(conn, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Accept a stream
	stream, err := session.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	autoConn(stream)
}

var autoHandler = gost.AutoHandler()

func autoServe(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		log.Println(conn)
		go autoConn(conn)
	}
	return nil
}

func autoConn(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("socks5server: recovered from panic:", r, conn)
		}
	}()
	autoHandler.Handle(conn)
	conn.Close()
}
