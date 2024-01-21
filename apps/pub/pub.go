package pub

import (
	"github.com/webteleport/ufo/apps/pub/handler"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	arg0 := Arg0(args, ".")
	h := handler.Handler(arg0)
	return wtf.Serve("https://k0s.io", h)
}
