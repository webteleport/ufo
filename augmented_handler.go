package ufo

import "net/http"

type AugmentedHandler interface {
	http.Handler
	PkgPath() string
}
