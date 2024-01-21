package handler

import (
	"fmt"
	"net/http"
)

var (
	username = "user"
	password = "pass"
)

func Handler() http.Handler {
	return http.HandlerFunc(handleRequest)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()

	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}

	if u != username {
		fmt.Printf("Username provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}

	if p != password {
		fmt.Printf("Password provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}

	fmt.Printf("Username: %s\n", u)
	fmt.Printf("Password: %s\n", p)

	w.WriteHeader(200)

	return
}
