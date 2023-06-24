package gos

import (
	"net/http"

	"github.com/webteleport/ufo/x"
	"github.com/webteleport/webteleport/ufo"
	"k0s.io/pkg/middleware"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	return ufo.Serve(Arg0(args, "https://ufo.k0s.io"), x.WellKnownHealthMiddleware(middleware.LoggingMiddleware(http.FileServer(http.Dir(".")))))
}
