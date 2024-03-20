package teleport

import (
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/ufo/apps/teleport/handler"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Arg1(args []string, fallback string) string {
	if len(args) > 1 {
		return args[1]
	}
	return fallback
}

func Run(args []string) error {
	addr := Arg0(args, apps.RELAY)
	upstream := Arg1(args, "https://k0s.io")
	return wtf.Serve(addr, handler.Handler(upstream))
}
