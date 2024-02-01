package basicauth

import (
	"context"
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps/basicauth/handler"
	"github.com/webteleport/webteleport"
)

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	return http.Serve(ln, handler.Handler())
}
