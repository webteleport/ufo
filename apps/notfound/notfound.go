package notfound

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	return wtf.Serve(apps.Arg0(args, apps.RELAY), utils.GinLoggerMiddleware(utils.HostNotFoundHandler()))
}
