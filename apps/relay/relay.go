package relay

import (
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"

	"github.com/caddyserver/certmagic"
	"github.com/webteleport/relay/session"
	"github.com/webteleport/relay/webteleport"
	"github.com/webteleport/ufo/apps/relay/envs"
	"github.com/webteleport/utils"
)

func OnDemandTLSConfig() (*tls.Config, error) {
	certmagic.DefaultACME.DisableHTTPChallenge = true
	tlsConfig, err := certmagic.TLS([]string{envs.HOST})
	if err != nil {
		return nil, err
	}
	tlsConfig.NextProtos = append([]string{"h2", "http/1.1"}, tlsConfig.NextProtos...)
	return tlsConfig, nil
}

func LazyTLSConfig(certFile, keyFile string) *tls.Config {
	GetCertificate := func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		// Always get latest localhost.crt and localhost.key
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}
	return &tls.Config{
		GetCertificate: GetCertificate,
	}
}

func listenHTTPS(handler http.Handler, errc chan error) {
	if envs.HTTPS_PORT == nil {
		return
	}
	slog.Info("listening on HTTPS https://" + envs.HOST + *envs.HTTPS_PORT)
	tlsConfig := LazyTLSConfig(envs.CERT, envs.KEY)
	ln, err := tls.Listen("tcp4", *envs.HTTPS_PORT, tlsConfig)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenHTTP(handler http.Handler, errc chan error) {
	if envs.HTTP_PORT == nil {
		return
	}
	slog.Info("listening on HTTP http://" + envs.HOST + *envs.HTTP_PORT)
	ln, err := net.Listen("tcp4", *envs.HTTP_PORT)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenTCPOnDemandTLS(handler http.Handler, errc chan error) {
	if envs.HTTPS_PORT == nil {
		return
	}
	slog.Info("listening on HTTPS https://" + envs.HOST + *envs.HTTPS_PORT + " w/ on demand tls")
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
	ln, err := tls.Listen("tcp4", *envs.HTTPS_PORT, tlsConfig)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenUDP(handler http.Handler, errc chan error) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT)
	tlsConfig := LazyTLSConfig(envs.CERT, envs.KEY)
	wts := webteleport.NewServerTLS(envs.HOST, envs.UDP_PORT, handler, tlsConfig)
	errc <- wts.ListenAndServe()
}

func listenUDPOnDemandTLS(handler http.Handler, errc chan error) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT + " w/ on demand tls")
	tlsConfig, err := OnDemandTLSConfig()
	if err != nil {
		errc <- err
		return
	}
	wts := webteleport.NewServerTLS(envs.HOST, envs.UDP_PORT, handler, tlsConfig)
	errc <- wts.ListenAndServe()
}

func listenAll(handler http.Handler) error {
	var errc chan error = make(chan error, 3)

	go listenUDP(handler, errc)
	go listenHTTP(handler, errc)
	go listenHTTPS(handler, errc)

	return <-errc
}

func Run([]string) error {
	var dsm http.Handler = session.DefaultSessionManager

	// Set the Alt-Svc header for UDP port discovery && http3 bootstrapping
	dsm = AltSvcMiddleware(dsm)
	dsm = utils.LoggingMiddleware(dsm)

	return listenAll(dsm)
}

func AltSvcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Alt-Svc", envs.ALT_SVC)
		next.ServeHTTP(w, r)
	})
}
