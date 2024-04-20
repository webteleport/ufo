package hello

import (
	"context"
	"io"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/webteleport"
)

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), apps.RELAY)
	if err != nil {
		return err
	}
	// log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, UFO!\n")
	})
	return http.Serve(ln, nil)
}
