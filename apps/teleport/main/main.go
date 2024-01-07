package main

import (
	"github.com/fermyon/spin/sdk/go/v2/http"
	"github.com/webteleport/ufo/apps/teleport/handler"
)

func init() {
	http.Handle(handler.Handler("https://example.com").ServeHTTP)
}

func main() {}
