package x

import "net/http"

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
// # Needs to have Jupyter extension installed
//
// - /login is used to check if server has authentication enabled
// - /hub/api is used to check if server is JupyterHub
//
// with authentication disabled, the /login endpoint won't respond with cors header
// so we need to manually add it to make vscode-web happy
//
// jupyter server doesn't have the /hub/api endpoint, so the same hack needs to be applied
func Jupyter(next http.Handler) http.Handler {
	jupyter := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/hub/api", "/login":
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(jupyter)
}
