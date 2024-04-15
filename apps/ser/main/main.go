package main

import (
	"github.com/fermyon/spin/sdk/go/v2/http"
	"github.com/webteleport/ufo/apps/ser/handler"
)

func init() {
	http.Handle(handler.Handler(".").ServeHTTP)
}

func main() {}
