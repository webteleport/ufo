package x

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func Formatter(writer io.Writer, params handlers.LogFormatterParams) {
	fmt.Fprintf(writer, "%s [%v] %s %s %s %s %d %d\n",
		params.Request.RemoteAddr,
		params.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
		params.Request.Host,
		params.Request.Method,
		params.Request.URL.RequestURI(),
		params.Request.Proto,
		params.StatusCode,
		params.Size,
	)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return handlers.CustomLoggingHandler(os.Stderr, next, Formatter)
}
