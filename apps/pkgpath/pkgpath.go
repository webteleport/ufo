package pkgpath

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps/pkgpath/mux"
	"github.com/webteleport/webteleport"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	wts := Arg0(args, "https://ufo.k0s.io")
	ln, err := webteleport.Listen(context.Background(), wts)
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ln.ClickableURL())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Host != ln.String() {
			mux.Mux.ServeHTTP(w, r)
			return
		}
		io.WriteString(w, "available handlers: "+mux.Mux.PkgPath())
	})
	return http.Serve(ln, nil)
}
