package handler

import (
	"io"
	"net/http"
)

// https://twitter.com/luogl/status/1748580566205927919
const IFRAME = `<iframe src="https://ip.skk.moe" style="width: 100%; height: 100%; border: 0"></iframe>`

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, IFRAME)
	})
}
