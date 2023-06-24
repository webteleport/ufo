// only works when using curl.
// modern browsers like chrome and firefox do not send the username/password.
package who

import (
	"io"
	"net/http"

	"github.com/webteleport/webteleport/ufo"
	"k0s.io/pkg/middleware"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func who(w http.ResponseWriter, r *http.Request) {
	// get request user password
	user, pass, ok := r.BasicAuth()
	io.WriteString(w, "user: "+user+"\n")
	io.WriteString(w, "pass: "+pass+"\n")
	if ok {
		io.WriteString(w, "ok: true\n")
	} else {
		io.WriteString(w, "ok: false\n")
	}
}

func Run(args []string) error {
	var h http.Handler
	h = http.HandlerFunc(who)
	h = middleware.LoggingMiddleware(h)
	return ufo.Serve(Arg0(args, "https://ufo.k0s.io"), h)
}
