package echo

import (
	"github.com/webteleport/ufo/apps/echo/handler"
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
	return wtf.Serve(Arg0(args, "https://ufo.k0s.io"), utils.GinLoggerMiddleware(handler.Handler()))
}
