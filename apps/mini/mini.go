package mini

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/webteleport/relay"
	"github.com/webteleport/ufo/apps/relay/envs"
	"github.com/webteleport/utils"
)

func listenHTTP(handler http.Handler) error {
	slog.Info("listening on HTTP http://" + envs.HOST + envs.HTTP_PORT)
	ln, err := net.Listen("tcp4", envs.HTTP_PORT)
	if err != nil {
		return err
	}
	return http.Serve(ln, handler)
}

func Run([]string) (err error) {
	s := relay.NewWSServer(envs.HOST, relay.NewSessionStore())

	var S http.Handler = s
	// Set the Alt-Svc header for UDP port discovery && http3 bootstrapping
	S = AltSvcMiddleware(S)
	S = utils.GinLoggerMiddleware(S)

	return listenHTTP(S)
}

func AltSvcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Alt-Svc", envs.ALT_SVC)
		next.ServeHTTP(w, r)
	})
}
