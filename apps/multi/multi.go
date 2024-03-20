package multi

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/webteleport"
)

var Map map[string]net.Listener = map[string]net.Listener{}

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
	ln, err := webteleport.Listen(context.Background(), apps.RELAY)
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	Map[webteleport.HumanURL(ln)] = ln
	err = http.Serve(ln, nil)
	log.Println(webteleport.HumanURL(ln), err)
	return err
}

func Run(args []string) error {
	wts := apps.Arg0(args, apps.RELAY)
	ln, err := webteleport.Listen(context.Background(), wts)
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", webteleport.ClickableURL(ln))
	Map[webteleport.HumanURL(ln)] = ln
	http.HandleFunc("/", index)
	return http.Serve(ln, nil)
}
