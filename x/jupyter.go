package x

import (
	"net/http"
	"strings"
)

// VSCode web edition makes cors request to /hub/api endpoint
// in order to determine if the backend api is a JupyterHub or JupyterLab/JupyterServer
//
// If the Access-Control-Allow-Origin header is absent, the request will not make through.
// Here we add this header to make VSCode web happy
//
// Tested on:
// - github.dev
// - vscode.dev
// - insider.vscode.dev
//
// Needs to have Jupyter extension installed
func Jupyter(next http.Handler) http.Handler {
	jupyter := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/hub/api") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(jupyter)
}
