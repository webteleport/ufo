package relay

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/caddyserver/certmagic"
	"github.com/quic-go/quic-go/http3"
	"github.com/webteleport/relay"
	"github.com/webteleport/ufo/apps/relay/envs"
	"github.com/webteleport/utils"
)

func OnDemandTLSConfig() (*tls.Config, error) {
	// if the decision function returns an error, a certificate
	// may not be obtained for that name at that time
	certmagic.Default.OnDemand = &certmagic.OnDemandConfig{
		DecisionFunc: func(_ctx context.Context, name string) error {
			if name != envs.HOST && name != fmt.Sprintf("*.%s", envs.HOST) {
				return fmt.Errorf("on-demand cert request denied for %s", name)
			}
			return nil
		},
	}
	// Because this convenience function returns only a TLS-enabled
	// listener and does not presume HTTP is also being served,
	// the HTTP challenge will be disabled. The package variable
	// Default is modified so that the HTTP challenge is disabled.
	certmagic.DefaultACME.DisableHTTPChallenge = true
	// tlsConfig := certmagic.Default.TLSConfig()
	tlsConfig, err := certmagic.TLS([]string{
		envs.HOST,
		fmt.Sprintf("*.%s", envs.HOST),
	})
	if err != nil {
		return nil, err
	}
	tlsConfig.NextProtos = append([]string{"h2", "http/1.1"}, tlsConfig.NextProtos...)
	return tlsConfig, nil
}

func useLocalTLS() bool {
	_, err := tls.LoadX509KeyPair(envs.CERT, envs.KEY)
	return err == nil
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
	slog.Info("listening on HTTP http://" + envs.HOST + envs.HTTP_PORT)
	ln, err := net.Listen("tcp4", envs.HTTP_PORT)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

func listenHTTPS(handler http.Handler, errc chan error, tlsConfig *tls.Config) {
	if envs.HTTPS_PORT == nil {
		return
	}
	slog.Info("listening on HTTPS https://" + envs.HOST + *envs.HTTPS_PORT)
	ln, err := tls.Listen("tcp4", *envs.HTTPS_PORT, tlsConfig)
	if err != nil {
		errc <- err
		return
	}
	errc <- http.Serve(ln, handler)
}

// TODO: refactor to http3.Serve(addr, t)
func listenWT(t http.Handler, errc chan error, tlsConfig *tls.Config) {
	slog.Info("listening on UDP https://" + envs.HOST + envs.UDP_PORT)
	h3s := http3.Server{
		Addr:            envs.UDP_PORT,
		TLSConfig:       tlsConfig,
		Handler:         t,
		EnableDatagrams: true,
	}

	errc <- h3s.ListenAndServe()
}

func listenAll(s http.Handler, t http.Handler, tlsConfig *tls.Config) error {
	var errc chan error = make(chan error, 3)

	go listenHTTP(s, errc)
	go listenHTTPS(s, errc, tlsConfig)
	go listenWT(t, errc, tlsConfig)

	return <-errc
}

func Run([]string) (err error) {
	var GlobalTLSConfig *tls.Config

	if useLocalTLS() {
		GlobalTLSConfig = LocalTLSConfig(envs.CERT, envs.KEY)
	} else {
		GlobalTLSConfig, err = OnDemandTLSConfig()
		if err != nil {
			slog.Warn("failed to get TLS config: ", err)
		}
	}

	store := relay.NewSessionStore()
	s := relay.NewWSServer(envs.HOST, store)
	t := relay.NewWTServer(envs.HOST, envs.UDP_PORT, store, GlobalTLSConfig)

	var S http.Handler = s
	S = AltSvcMiddleware(S)
	S = utils.GinLoggerMiddleware(S)

	var T http.Handler = t
	// Set the Alt-Svc header for UDP port discovery && http3 bootstrapping
	T = AltSvcMiddleware(T)
	T = utils.GinLoggerMiddleware(T)

	return listenAll(S, T, GlobalTLSConfig)
}

func AltSvcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Alt-Svc", envs.ALT_SVC)
		next.ServeHTTP(w, r)
	})
}
