package cookies

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/btwiuse/rng"
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport"
)

func Run(args []string) error {
	ln, err := webteleport.Listen(context.Background(), apps.Arg0(args, apps.RELAY))
	if err != nil {
		return err
	}
	tkn := rng.NewUUID()
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
	// log.Println("listening on", webteleport.ClickableURL(ln))
	// log.Println("access link", webteleport.ClickableURL(ln)+secretPath)
	return http.Serve(ln, utils.GinLoggerMiddleware(http.DefaultServeMux))
}
