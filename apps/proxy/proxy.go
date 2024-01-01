package proxy

import (
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

func info(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.Path, r.Proto)
			log.Println(r.Host)
			log.Println(r.URL.Query())
			log.Println(r.Header)
			next.ServeHTTP(w, r)
		},
	)
}

func Run(args []string) error {
	var handler http.Handler
	handler = http.DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "OK", 200)
	})
	handler = info(utils.GinLoggerMiddleware(handler))
	arg0 := Arg0(args, "https://ufo.k0s.io")
	if arg0 == "local" {
		port := utils.EnvPort(":8000")
		log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
		return http.ListenAndServe(port, handler)
	}
	return ufo.Serve(arg0, handler)
}
