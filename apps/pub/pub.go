package pub

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/webteleport/ufo/apps/pub/handler"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func partitionIntoPairs(s []string) [][2]string {
	var pairs [][2]string

	for i := 0; i < len(s); i += 2 {
		if i+1 < len(s) {
			pairs = append(pairs, [2]string{s[i], s[i+1]})
		} else {
			pairs = append(pairs, [2]string{s[i]})
		}
	}

	return pairs
}

func Run(args []string) error {
	mux := http.NewServeMux()

	arg0 := Arg0(args, ".")

	if arg0 == "--" {
		pairs := partitionIntoPairs(args[1:])
		for _, pair := range pairs {
			slog.Info(fmt.Sprintf("publishing: %s -> %s", pair[0], pair[1]))
			mux.Handle(pair[0], http.StripPrefix(strings.TrimSuffix(pair[0], "/"), handler.Handler(pair[1])))
		}
	} else {
		slog.Info(fmt.Sprintf("publishing: %s", arg0))
		mux.Handle("/", handler.Handler(arg0))
	}

	return wtf.Serve("https://k0s.io", mux)
}
