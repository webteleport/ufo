package ip

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/ip/handler"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	return wtf.Serve(Arg0(args, apps.RELAY), utils.GinLoggerMiddleware(handler.Handler()))
}
