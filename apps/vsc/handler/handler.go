package handler

import (
	"io"
	"net/http"

	"github.com/btwiuse/pretty"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, pretty.JSONStringLine(r.Header))
	})
}
