package x

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(addr string) http.Handler {
	rpURL, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	rp := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(rpURL)
			r.SetXForwarded()
		},
	}
	return rp
}
