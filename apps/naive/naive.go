package naive

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"github.com/smarty/cproxy/v2"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport"
)

// local listener works
//
//	https_proxy=http://a:b@localhost:8080 curl https://google.com -v
//
// while the remote listener doesn't
//
//	http_proxy=https://naive.remotehost:pass@remotehost:8444 curl http://google.com -v
func Run([]string) error {
	ln1, err := net.Listen("tcp", ":8080")
	log.Println("🛸 listening on", ":8080")
	ln2, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io:8444/naive?naive=1")
	// log.Println("🛸 listening on", webteleport.ClickableURL(ln2))
	if err != nil {
		return err
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	// Create a BasicAuth middleware with the provided credentials
	basic := auth.BasicConnect(
		"Restricted",
		func(user, pass string) bool {
			// ok := user != "" && pass != ""
			// return true || ok
			return true
		},
	)
	c := cproxy.New()

	// Use the BasicAuth middleware with the proxy server
	proxy.OnRequest().HandleConnect(basic)
	a := &A{proxy, c}
	b := utils.GinLoggerMiddleware(a)
	go http.Serve(ln1, b)
	return http.Serve(ln2, b)
}

type A struct {
	proxy *goproxy.ProxyHttpServer
	c     http.Handler
}

func (a *A) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	// a.c.ServeHTTP(w, r)
	a.proxy.ServeHTTP(w, r)
	println()
}
