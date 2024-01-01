package mux

import (
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/webteleport/utils"
)

var _ utils.AugmentedHandler = (*ServeMux)(nil)

type ServeMux struct {
	*http.ServeMux
}

func (sm *ServeMux) PkgPath() string {
	return reflect.TypeOf(*sm).PkgPath()
}

var Mux *ServeMux = NewServeMux()

func NewServeMux() *ServeMux {
	mux := http.NewServeMux()
	wrapped := &ServeMux{ServeMux: mux}
	pkgpath := wrapped.PkgPath()
	// log.Println(pkgpath)
	wrapped.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(pkgpath)
		io.WriteString(w, pkgpath)
	})
	return wrapped
}
