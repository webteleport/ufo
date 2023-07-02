package login

import (
	"net/http"
	"net/url"

	"github.com/webteleport/auth"
	"github.com/webteleport/webteleport/ufo"
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
	lm := &auth.LoginMiddleware{
		Password: u.Fragment,
	}
	// support listing files under cwd, but not actual file is served
	// making this perfect for demonstration purpose without compromising security
	cwd := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(r.URL.Path, http.FileServer(http.Dir("."))).ServeHTTP(w, r)
	})
	return ufo.Serve(stationURL, middleware.LoggingMiddleware(lm.Wrap(cwd)))
}
