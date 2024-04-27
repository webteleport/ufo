package logbody

import (
	"io"
	"net/http"
	"os"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	var handler http.Handler
	handler = http.DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(os.Stderr, r.Body)
		http.Error(w, "OK", 200)
	})
	handler = utils.GinLoggerMiddleware(handler)
	arg0 := apps.Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler)
}
