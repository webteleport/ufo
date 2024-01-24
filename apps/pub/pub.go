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

func registerRules(mux *http.ServeMux, args []string) {
	if len(args)%2 != 0 {
		slog.Warn("uneven number of args passed as rules, dropping the last")
	}
	pairs := partitionIntoPairs(args[1:])
	for _, pair := range pairs {
		slog.Info(fmt.Sprintf("publishing: %s -> %s", pair[0], pair[1]))
		mux.Handle(pair[0], http.StripPrefix(strings.TrimSuffix(pair[0], "/"), handler.Handler(pair[1])))
	}
}

func Run(args []string) error {
	mux := http.NewServeMux()

	arg0 := Arg0(args, ".")

	switch {
	case arg0 == "--":
		registerRules(mux, args[1:])
	case len(args) >= 2:
		slog.Warn(fmt.Sprintf("add -- before rules to remove ambiguity: pub -- %s", strings.Join(args, " ")))
		registerRules(mux, args)
	default:
		slog.Info(fmt.Sprintf("publishing: %s", arg0))
		mux.Handle("/", handler.Handler(arg0))
	}

	return wtf.Serve("https://k0s.io", mux)
}
