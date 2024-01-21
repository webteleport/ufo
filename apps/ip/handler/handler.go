package handler

import (
	"io"
	"net/http"
)

const IFRAME = `<iframe src="https://ip.skk.moe/simple-dark" style="width: 100%; border: 0"></iframe>`

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, IFRAME)
	})
}
