package hdr

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/hdr/handler"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	return wtf.Serve(apps.Arg0(args, apps.RELAY), utils.GinLoggerMiddleware(handler.Handler()))
}
