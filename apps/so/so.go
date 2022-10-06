// it doesn't work:
//
// $ export HTTP_PROXY=https://26.ufo.k0s.io HTTPS_PROXY=https://26.ufo.k0s.io
//
// $ curl https://google.com
// curl: (56) Received HTTP code 404 from proxy after CONNECT
//
// for HTTP_PROXY r.Host = google.com r.Method = Get
// for HTTPS_PROXY r.Host = google.com r.Method = Connect
// they will not be supported by ufo server
package so

import (
	"context"
	"log"
	"net"

	"github.com/ginuerzh/gost"
	"github.com/webteleport/webteleport"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	ln, err := webteleport.Listen(context.Background(), Arg0(args, "https://ufo.k0s.io"))
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ln.ClickableURL())
	return autoServe(ln)
}

var autoHandler = gost.AutoHandler()

func autoServe(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		log.Println(conn)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("socks5server: recovered from panic:", r, conn)
				}
			}()
			autoHandler.Handle(conn)
			conn.Close()
		}()
	}
	return nil
}
