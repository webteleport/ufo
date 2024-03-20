package proxy

import (
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

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
	arg0 := apps.Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler)
}
