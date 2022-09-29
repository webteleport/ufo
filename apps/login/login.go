package login

import (
	"net/http"
	"net/url"

	"github.com/webteleport/webteleport"
	"k0s.io/pkg/middleware"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	stationURL := Arg0(args, "https://ufo.k0s.io")
	u, err := url.Parse(stationURL)
	if err != nil {
		return nil
	}
	lm := &LoginMiddleware{
		Password: u.Fragment,
	}
	pwd := http.FileServer(http.Dir("."))
	return webteleport.Serve(stationURL, middleware.LoggingMiddleware(lm.Wrap(pwd)))
}
