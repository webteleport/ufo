package metrics

import (
	"net/http"

	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport/ufo"
	"k0s.io/pkg/exporter"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

var Handler http.Handler = utils.LoggingMiddleware(utils.GzipMiddleware(exporter.NewHandler()))

func Run(args []string) error {
	return ufo.Serve(Arg0(args, "https://metrics.k0s.io"), Handler)
}
