package ser

import (
	"github.com/webteleport/ufo/apps/ser/handler"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	arg0 := Arg0(args, "https://ufo.k0s.io")
	h := handler.Handler(".")
	return wtf.Serve(arg0, h)
}
