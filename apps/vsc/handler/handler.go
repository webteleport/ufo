package handler

import (
	"io"
	"net/http"

	"github.com/btwiuse/pretty"
)

const QUALITYLESS_PRODUCT_NAME = "Code"
const QUALITYLESS_SERVER_NAME = QUALITYLESS_PRODUCT_NAME + " Server"

func waitForDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	html := "The latest version of the " + QUALITYLESS_SERVER_NAME + " is downloading, please wait a moment...<script>setTimeout(()=>location.reload(),1500)</script>"
	w.Write([]byte(html))
}

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		waitForDownload(w, r)
		return
		io.WriteString(w, pretty.JSONStringLine(r.Header))
	})
}
