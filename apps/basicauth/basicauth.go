package basicauth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/webteleport/webteleport"
)

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ln.ClickableURL())
	http.HandleFunc("/", handleRequest)
	return http.Serve(ln, nil)
}

var (
	username = "user"
	password = "pass"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}
	if u != username {
		fmt.Printf("Username provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}
	if p != password {
		fmt.Printf("Password provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("Username: %s\n", u)
	fmt.Printf("Password: %s\n", p)
	w.WriteHeader(200)
	return
}
