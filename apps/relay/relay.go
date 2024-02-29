package relay

import (
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"

	"github.com/caddyserver/certmagic"
	"github.com/webteleport/relay"
	"github.com/webteleport/relay/manager"
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

func useLocalTLS() bool {
	_, err := tls.LoadX509KeyPair(envs.CERT, envs.KEY)
	return err != nil
}

func LocalTLSConfig(certFile, keyFile string) *tls.Config {
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

func listenHTTPSLocalTLS(handler http.Handler, errc chan error) {
	if envs.HTTPS_PORT == nil {
		return
	}
	slog.Info("listening on HTTPS https://" + envs.HOST + *envs.HTTPS_PORT)
	tlsConfig := LocalTLSConfig(envs.CERT, envs.KEY)
	ln, err := tls.Listen("tcp4", *envs.HTTPS_PORT, tlsConfig)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenHTTPSOnDemandTLS(handler http.Handler, errc chan error) {
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

func listenUDPLocalTLS(handler http.Handler, errc chan error) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT)
	tlsConfig := LocalTLSConfig(envs.CERT, envs.KEY)
	r := relay.New(envs.HOST, envs.UDP_PORT, handler, tlsConfig)
	errc <- r.ListenAndServe()
}

func listenUDPOnDemandTLS(handler http.Handler, errc chan error) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT + " w/ on demand tls")
	tlsConfig, err := OnDemandTLSConfig()
	if err != nil {
		errc <- err
		return
	}
	r := relay.New(envs.HOST, envs.UDP_PORT, handler, tlsConfig)
	errc <- r.ListenAndServe()
}

func listenAll(handler http.Handler) error {
	var errc chan error = make(chan error, 3)

	go listenHTTP(handler, errc)
	if useLocalTLS() {
		go listenUDPLocalTLS(handler, errc)
		go listenHTTPSLocalTLS(handler, errc)
	} else {
		go listenUDPOnDemandTLS(handler, errc)
		go listenHTTPSOnDemandTLS(handler, errc)
	}

	return <-errc
}

func Run([]string) error {
	var dsm http.Handler = manager.DefaultSessionManager

	// Set the Alt-Svc header for UDP port discovery && http3 bootstrapping
	dsm = AltSvcMiddleware(dsm)
	dsm = utils.GinLoggerMiddleware(dsm)

	return listenAll(dsm)
}

func AltSvcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Alt-Svc", envs.ALT_SVC)
		next.ServeHTTP(w, r)
	})
}
