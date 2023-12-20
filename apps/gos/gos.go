package gos

import (
	"fmt"
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
	err := ufo.Serve(Arg0(args, "https://ufo.k0s.io"), middleware.LoggingMiddleware(x.WellKnownHealthMiddleware(middleware.GzipMiddleware(http.FileServer(http.Dir("."))))))
	if err != nil {
		return fmt.Errorf("ufo: %w", err)
	}
	return nil
}
