package gopilot

import (
	"net/http"

	"github.com/chadgpt/gopilot"
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/wbdv"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	mux := http.NewServeMux()
	wbdvHandler := wbdv.Handler(".")
	mux.Handle("/chatgpt-next-web", wbdvHandler)
	mux.Handle("/chatgpt-next-web/", wbdvHandler)
	mux.Handle("/", gopilot.Handler())
	handler := utils.GzipMiddleware(mux)
	handler = utils.GinLoggerMiddleware(handler)
	handler = utils.AllowAllCorsMiddleware(handler)
	arg0 := apps.Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler)
}
