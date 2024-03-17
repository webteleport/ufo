package handler

import (
	"net/http"

	"github.com/webteleport/utils"
)

const QUALITYLESS_PRODUCT_NAME = "Code"
const QUALITYLESS_SERVER_NAME = QUALITYLESS_PRODUCT_NAME + " Server"

func waitForDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	html := "The latest version of the " + QUALITYLESS_SERVER_NAME + " is downloading, please wait a moment...<script>setTimeout(()=>location.reload(),1500)</script>"
	w.Write([]byte(html))
}

func Handler(addr string) http.Handler {
	handler := utils.ReverseProxy(addr)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
}
