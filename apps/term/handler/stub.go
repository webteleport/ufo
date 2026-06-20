//go:build js || wasip1

package handler

import (
	"errors"
	"net/http"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, errors.New("term: not supported on this platform").Error(), http.StatusNotImplemented)
	})
}
