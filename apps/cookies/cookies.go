package cookies

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/webteleport/webteleport"
	"k0s.io/pkg/middleware"
	"k0s.io/pkg/uuid"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	ln, err := webteleport.Listen(context.Background(), Arg0(args, "https://ufo.k0s.io"))
	if err != nil {
		return err
	}
	tkn := uuid.New()
	secretPath := "/.secret-path/" + tkn
	http.HandleFunc(secretPath, func(w http.ResponseWriter, r *http.Request) {
		cookies := fmt.Sprintf(`WebTeleportAccessToken="%s"; Path=/; Max-Age=2592000; HttpOnly; Domain=%s`, tkn, r.Host)
		w.Header().Set("Set-Cookie", cookies)
		http.Redirect(w, r, "/", 302)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wtat, err := r.Cookie("WebTeleportAccessToken")
		if err != nil {
			http.Error(w, "🛸"+http.StatusText(401)+" "+err.Error(), 401)
			return
		}
		if wtat.Value != tkn {
			http.Error(w, "🛸"+http.StatusText(401), 401)
			return
		}
		io.WriteString(w, "🛸"+http.StatusText(200))
	})
	log.Println("listening on", ln.ClickableURL())
	log.Println("access link", ln.ClickableURL()+secretPath)
	return http.Serve(ln, middleware.LoggingMiddleware(http.DefaultServeMux))
}
