package gos

import (
	"fmt"
	"log"
	"net/http"

	"github.com/webteleport/utils"
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
	handler := middleware.LoggingMiddleware(utils.WellKnownHealthMiddleware(middleware.GzipMiddleware(http.FileServer(http.Dir(".")))))
	arg0 := Arg0(args, "https://ufo.k0s.io")
	if arg0 == "local" {
		port := utils.EnvPort(":8000")
		log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
		return http.ListenAndServe(port, handler)
	}
	return ufo.Serve(arg0, handler)
}
