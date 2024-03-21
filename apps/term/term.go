package term

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/term/handler"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	arg0 := apps.Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler.Handler())
}
