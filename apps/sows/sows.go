// $ websocat tcp-listen:0.0.0.0:3000 --binary wss://35.ufo.k0s.io -E
// $ export HTTP_PROXY=http://127.0.0.1:3000 HTTPS_PROXY=http://127.0.0.1:3000

// $ websocat tcp-listen:0.0.0.0:1080 --binary wss://sows.ufo.k0s.io -E
// $ export HTTP_PROXY=http://127.0.0.1:1080 HTTPS_PROXY=http://127.0.0.1:1080

// $ curl -v http://google.com
// $ curl -v https://google.com
package sows

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/btwiuse/gost"
	"github.com/webteleport/ufo/x"
	"github.com/webteleport/webteleport"

	"k0s.io/pkg/middleware"
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

func Run(args []string) error {
	ln, err := webteleport.Listen(context.Background(), Arg0(args, "https://ufo.k0s.io"))
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", wsfy(ln.Network())+"://"+ln.String())

	return http.Serve(ln, middleware.LoggingMiddleware(x.WellKnownHealthMiddleware(&auto{})))
}

type auto struct{}

func (a *auto) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wrap.Wrconn(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	autoConn(conn)
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
