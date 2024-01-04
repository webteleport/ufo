package echo

import (
	echo "github.com/jpillora/go-echo-server/handler"
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
	e := echo.New(echo.Config{})
	e = utils.GinLoggerMiddleware(e)
	return wtf.Serve(Arg0(args, "https://ufo.k0s.io"), e)
}
