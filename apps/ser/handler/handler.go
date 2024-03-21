package handler

import (
	"expvar"
	"net/http"

	"github.com/btwiuse/better"
	"github.com/webteleport/utils"
)

func Handler(s string) http.Handler {
	handler := better.FileServer(http.Dir(s))
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)
	return mux
}

func RootHandler() http.Handler {
	handler := Handler("/")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
