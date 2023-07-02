package x

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(addr string) http.Handler {
	upstream, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	rewrite := func(r *httputil.ProxyRequest) {
		r.SetURL(upstream)
		r.SetXForwarded()

		// use client credentials if available
		if r.In.Header.Get("Authorization") != "" {
			return
		}

		// otherwise use credentials from upstream
		if upstream.User != nil {
			user := upstream.User.Username()
			pass, _ := upstream.User.Password()
			r.Out.SetBasicAuth(user, pass)
		}
	}
	rp := &httputil.ReverseProxy{
		Rewrite: rewrite,
	}
	return rp
}
