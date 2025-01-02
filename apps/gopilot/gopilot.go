package gopilot

import (
	"github.com/chadgpt/gopilot"
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	handler := gopilot.Handler()
	handler = utils.GzipMiddleware(handler)
	handler = utils.AllowAllCorsMiddleware(handler)
	arg0 := apps.Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler)
}
