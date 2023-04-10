package ufo

import "net/http"

func WellKnownHealthMiddleware(next http.Handler) http.Handler {
	health := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/.well-known/health" {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	return http.HandlerFunc(health)
}
