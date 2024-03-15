package handler

import (
	"io"
	"net/http"
)

// https://twitter.com/luogl/status/1748580566205927919
const IFRAME = `<iframe src="https://ip.skk.moe" style="width: 100%; height: 100%; border: 0"></iframe>`

func OldHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, IFRAME)
	})
}

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
