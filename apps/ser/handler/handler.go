package handler

import (
	"expvar"
	"net/http"

	"github.com/webteleport/utils"
)

func Handler(s string) http.Handler {
	handler := http.FileServer(http.Dir(s))
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)
	return mux
}
