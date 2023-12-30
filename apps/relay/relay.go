package relay

import (
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"

	"github.com/caddyserver/certmagic"
	"github.com/webteleport/server/envs"
	"github.com/webteleport/server/session"
	"github.com/webteleport/server/webteleport"
	"github.com/webteleport/ufo/x"
)

func listenTCP(handler http.Handler, errc chan error) {
	slog.Info("listening on TCP http://" + envs.HOST + envs.TCP_PORT)
	ln, err := net.Listen("tcp4", envs.TCP_PORT)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenTCPOnDemandTLS(handler http.Handler, errc chan error) {
	slog.Info("listening on TCP https://" + envs.HOST + envs.TCP_PORT + " w/ on demand tls")
	// Because this convenience function returns only a TLS-enabled
	// listener and does not presume HTTP is also being served,
	// the HTTP challenge will be disabled. The package variable
	// Default is modified so that the HTTP challenge is disabled.
	certmagic.DefaultACME.DisableHTTPChallenge = true
	// tlsConfig := certmagic.Default.TLSConfig()
	tlsConfig, err := certmagic.TLS([]string{envs.HOST})
	if err != nil {
		errc <- err
		return
	}
	tlsConfig.NextProtos = append([]string{"h2", "http/1.1"}, tlsConfig.NextProtos...)
	ln, err := tls.Listen("tcp4", envs.TCP_PORT, tlsConfig)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenUDP(handler http.Handler, errc chan error) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT)
	wts := webteleport.NewServerTLS(handler, envs.CERT, envs.KEY)
	errc <- wts.ListenAndServe()
}

func listenUDPOnDemandTLS(handler http.Handler, errc chan error) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT + " w/ on demand tls")
	wts, err := webteleport.NewServerTLSOnDemand(handler)
	if err != nil {
		errc <- err
		return
	}
	errc <- wts.ListenAndServe()
}

func listenAll(handler http.Handler) error {
	var errc chan error = make(chan error, 2)

	go listenTCPOnDemandTLS(handler, errc)
	go listenUDPOnDemandTLS(handler, errc)

	return <-errc
}

func Run([]string) error {
	var dsm http.Handler = session.DefaultSessionManager

	dsm = x.WellKnownHealthMiddleware(dsm)
	dsm = x.LoggingMiddleware(dsm)

	return listenAll(dsm)
}
