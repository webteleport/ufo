package grafana

import (
	"os"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/teleport/handler"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	os.Setenv("PASS_HOST", "1")
	addr := apps.Arg0(args, apps.RELAY)
	upstream := apps.Arg1(args, "https://k0s.io")
	return wtf.Serve(addr, handler.Handler(upstream))
}
