package login

import (
	"net/http"
	"net/url"

	"github.com/webteleport/auth"
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	stationURL := apps.Arg0(args, apps.RELAY)
	u, err := url.Parse(stationURL)
	if err != nil {
		return err
	}
	// support listing files under cwd, but not actual file is served
	// making this perfect for demonstration purpose without compromising security
	cwd := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(r.URL.Path, http.FileServer(http.Dir("."))).ServeHTTP(w, r)
	})
	return wtf.Serve(stationURL, utils.GinLoggerMiddleware(auth.WithPassword(cwd, u.Fragment)))
}
