package pkgpath

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/pkgpath/mux"
	"github.com/webteleport/webteleport"
)

func Run(args []string) error {
	wts := apps.Arg0(args, apps.RELAY)
	ln, err := webteleport.Listen(context.Background(), wts)
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Host != ln.Addr().String() {
			mux.Mux.ServeHTTP(w, r)
			return
		}
		io.WriteString(w, "available handlers: "+mux.Mux.PkgPath())
	})
	return http.Serve(ln, nil)
}
