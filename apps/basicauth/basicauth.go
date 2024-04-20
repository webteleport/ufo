package basicauth

import (
	"context"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/basicauth/handler"
	"github.com/webteleport/webteleport"
)

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), apps.RELAY)
	if err != nil {
		return err
	}
	// log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	return http.Serve(ln, handler.Handler())
}
