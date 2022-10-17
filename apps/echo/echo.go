package echo

import (
	echo "github.com/jpillora/go-echo-server/handler"
	"github.com/webteleport/ufo"
	"k0s.io/pkg/middleware"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	e := echo.New(echo.Config{})
	e = middleware.LoggingMiddleware(e)
	return ufo.Serve(Arg0(args, "https://ufo.k0s.io"), e)
}
