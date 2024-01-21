package handler

import (
	"expvar"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/webteleport/utils"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// File exists but other error occurred
	return false
}

func isPort(s string) bool {
	match, _ := regexp.MatchString(`^:\d{1,5}$`, s)
	return match
}

func isHostPort(s string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9.-]+:\d{1,5}$`, s)
	return match
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	return err == nil
}

func Handler(s string) (handler http.Handler) {
	switch {
	case pathExists(s):
		slog.Info(fmt.Sprintf("publishing path: %s", s))
		handler = http.FileServer(http.Dir(s))
	case isPort(s):
		slog.Info(fmt.Sprintf("publishing port: %s", s))
		handler = utils.ReverseProxy(s)
	case isHostPort(s):
		slog.Info(fmt.Sprintf("publishing http: %s", s))
		handler = utils.ReverseProxy(s)
	case isValidURL(s):
		slog.Info(fmt.Sprintf("publishing url: %s", s))
		handler = utils.ReverseProxy(s)
	default:
		log.Fatalln("unrecognized pattern:", s)
	}
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)
	return mux
}
