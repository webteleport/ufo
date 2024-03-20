package intercept

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	addr := apps.Arg0(args, apps.RELAY)
	upstream := apps.Arg1(args, "https://example.com")
	handler := utils.ReverseProxy(upstream)
	handler = utils.GinLoggerMiddleware(handler)
	handler = utils.InterceptMiddleware(handler)
	return wtf.Serve(addr, handler)
}
