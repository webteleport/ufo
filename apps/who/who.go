// only works when using curl.
// modern browsers like chrome and firefox do not send the username/password.
package who

import (
	"io"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

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
	h = utils.GinLoggerMiddleware(h)
	return wtf.Serve(apps.Arg0(args, apps.RELAY), h)
}
