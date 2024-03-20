package ser

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/ser/handler"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	arg0 := apps.Arg0(args, apps.RELAY)
	h := handler.Handler(".")
	return wtf.Serve(arg0, h)
}
