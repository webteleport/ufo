package handler

import (
	"io"
	"os"
	"net/http"

	"github.com/btwiuse/pretty"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(io.MultiWriter(w, os.Stderr), pretty.JSONStringLine(r.Header))
	})
}
