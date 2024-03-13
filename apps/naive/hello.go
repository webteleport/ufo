package naive

import (
	"context"
	"log"
	"net/http"

	"github.com/webteleport/webteleport"
)

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io?naive=1")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ConnectHandler(w, r)
	})
	return http.Serve(ln, nil)
}
