package wagi

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/webteleport/relay"
	"github.com/webteleport/ufo/apps/relay/envs"
	"github.com/webteleport/utils"
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
	s := relay.DefaultWSServer(envs.HOST)
	s.Use(utils.GinLoggerMiddleware)

	return listenHTTP(s)
}
