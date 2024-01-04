package metrics

import (
	"net/http"

	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
	"k0s.io/pkg/exporter"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

var Handler http.Handler = utils.GinLoggerMiddleware(utils.GzipMiddleware(exporter.NewHandler()))

func Run(args []string) error {
	return wtf.Serve(Arg0(args, "https://metrics.k0s.io"), Handler)
}
