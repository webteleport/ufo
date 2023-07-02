package dl

import (
	"net/http"
	"os"

	"github.com/webteleport/webteleport/ufo"
	"k0s.io/pkg/middleware"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func binHandler() http.Handler {
	exe, err := os.Executable()
	if err != nil {
		return http.NotFoundHandler()
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, exe)
	})
}

func Run(args []string) error {
	return ufo.Serve(Arg0(args, "https://ufo.k0s.io"), middleware.LoggingMiddleware(middleware.GzipMiddleware(binHandler())))
}