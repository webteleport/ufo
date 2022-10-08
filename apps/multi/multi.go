package multi

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/webteleport/webteleport"
)

var Map map[string]*webteleport.Listener = map[string]*webteleport.Listener{}

func index(w http.ResponseWriter, r *http.Request) {
	domain := strings.TrimPrefix(r.URL.Path, "/")
	key := "https://" + domain
	if ln, ok := Map[key]; ok {
		delete(Map, key)
		ln.Close()
		return
	}
	for k, _ := range Map {
		io.WriteString(w, k+"\n")
	}
	go newRoute()
}

func newRoute() error {
	ln, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ln.ClickableURL())
	Map[ln.HumanURL()] = ln
	err = http.Serve(ln, nil)
	log.Println(ln.HumanURL(), err)
	return err
}

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ln.ClickableURL())
	Map[ln.HumanURL()] = ln
	http.HandleFunc("/", index)
	return http.Serve(ln, nil)
}
