package naive

import (
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	// Create a BasicAuth middleware with the provided credentials
	basic := auth.BasicConnect(
		"Restricted",
		func(user, pass string) bool {
			ok := user != "" && pass != ""
			return ok
		},
	)

	// Use the BasicAuth middleware with the proxy server
	proxy.OnRequest().HandleConnect(basic)

	proxy.ServeHTTP(w, r)
}
