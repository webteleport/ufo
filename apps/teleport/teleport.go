package teleport

import (
	"github.com/webteleport/ufo"
	"k0s.io/pkg/middleware"
	"k0s.io/pkg/reverseproxy"
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
	addr := Arg0(args, "https://ufo.k0s.io")
	upstream := Arg1(args, "https://k0s.io")
	handler := reverseproxy.Handler(upstream)
	handler = middleware.AllowAllCorsMiddleware(handler)
	handler = middleware.LoggingMiddleware(handler)
	return ufo.Serve(addr, handler)
}
