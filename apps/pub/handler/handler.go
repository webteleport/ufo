package handler

import (
	"expvar"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/webteleport/utils"
)

// passing symlink also returns true
func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

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

func serveFile(s string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s)
	})
}

func Handler(s string) (handler http.Handler) {
	switch {
	case isFile(s):
		handler = serveFile(s)
	case pathExists(s):
		handler = http.FileServer(http.Dir(s))
	case isPort(s):
		handler = utils.ReverseProxy(s)
	case isHostPort(s):
		handler = utils.ReverseProxy(s)
	case isValidURL(s):
		handler = utils.ReverseProxy(s)
	default:
		handler = utils.ReverseProxy(s)
	}
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)
	return mux
}
