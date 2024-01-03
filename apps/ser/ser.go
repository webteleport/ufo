package ser

import (
	"expvar"
	"fmt"
	"log"
	"net/http"

	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport/ufo"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	handler := http.FileServer(http.Dir("."))
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	go collectMemstats()
	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)
	arg0 := Arg0(args, "https://ufo.k0s.io")
	if arg0 == "local" {
		port := utils.EnvPort(":8000")
		log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
		return http.ListenAndServe(port, mux)
	}
	return ufo.Serve(arg0, mux)
}
