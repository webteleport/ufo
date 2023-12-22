package server

import (
	"log"
	"net"
	"net/http"

	"github.com/webteleport/server/envs"
	"github.com/webteleport/server/session"
	"github.com/webteleport/server/webteleport"
	"github.com/webteleport/ufo/x"
)

func listenTCP(handler http.Handler, errc chan error) {
	log.Println("listening on TCP http://" + envs.HOST + envs.PORT)
	ln, err := net.Listen("tcp4", envs.PORT)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenUDP(handler http.Handler, errc chan error) {
	log.Println("listening on UDP https://" + envs.HOST + envs.PORT)
	wts := webteleport.NewServer(handler)
	errc <- wts.ListenAndServeTLS(envs.CERT, envs.KEY)
}

func listenAll(handler http.Handler) error {
	var errc chan error = make(chan error, 2)

	go listenTCP(handler, errc)
	go listenUDP(handler, errc)

	return <-errc
}

func Run([]string) error {
	var dsm http.Handler = session.DefaultSessionManager

	dsm = x.WellKnownHealthMiddleware(dsm)
	dsm = x.LoggingMiddleware(dsm)

	return listenAll(dsm)
}
