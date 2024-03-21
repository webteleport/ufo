package agent

import (
	"context"
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps"
	ser "github.com/webteleport/ufo/apps/ser/handler"
	term "github.com/webteleport/ufo/apps/term/handler"
	"github.com/webteleport/webteleport"
)

var Map = map[string]http.Handler{
	"/term": term.Handler(),
	"/fs":   ser.RootHandler(),
}

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), apps.RELAY)
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	mux := http.NewServeMux()
	for prefix, handler := range Map {
		mux.Handle(prefix+"/", http.StripPrefix(prefix, handler))
	}
	return http.Serve(ln, mux)
}
