package handler

import (
	"net/http"

	"github.com/webteleport/utils"
)

func Handler(upstream string) http.Handler {
	handler := utils.ReverseProxy(upstream)
	handler = utils.Jupyter(handler)
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	return handler
}
