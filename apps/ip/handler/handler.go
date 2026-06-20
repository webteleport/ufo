package handler

import (
	"io"
	"net/http"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(w, GetClientIP(r))
	})
}

func GetClientIP(r *http.Request) (clientIP string) {
	// Retrieve the client IP address from the request headers
	for _, x := range []string{
		r.Header.Get("X-Envoy-External-Address"),
		r.Header.Get("X-Real-IP"),
		r.Header.Get("X-Forwarded-For"),
		r.RemoteAddr,
	} {
		if x != "" {
			clientIP = x
			break
		}
	}
	return
}
