package mini

import (
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/webteleport/relay"
	"github.com/webteleport/ufo/apps/relay/envs"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport/transport/websocket"
	"github.com/webteleport/wtf"
)

func listenHTTP(handler http.Handler) error {
	slog.Info("listening on HTTP http://0.0.0.0"+envs.HTTP_PORT, "HOST", envs.HOST)
	ln, err := net.Listen("tcp", envs.HTTP_PORT)
	if err != nil {
		return err
	}
	return http.Serve(ln, handler)
}

func Run([]string) (err error) {
	store := relay.NewSessionStore()
	if os.Getenv("LOGGIN") != "" {
		store.Use(utils.GinLoggerMiddleware)
		store.Use(AltSvcMiddleware)
	}
	s := relay.NewWSServer(envs.HOST, store)

	extra := os.Getenv("EXTRA")
	if extra != "" {
		upgrader := &websocket.Upgrader{
			HOST: envs.HOST,
		}
		go wtf.Serve(extra, upgrader)
		go store.Subscribe(upgrader)
	}

	return listenHTTP(s)
}

func AltSvcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Alt-Svc", envs.ALT_SVC)
		next.ServeHTTP(w, r)
	})
}
