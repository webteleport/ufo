package wbdv

import (
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
	"golang.org/x/net/webdav"
)

func Run(args []string) error {
	handler := Handler(".")
	handler = utils.GzipMiddleware(handler)
	handler = utils.GinLoggerMiddleware(handler)
	handler = utils.AllowAllCorsMiddleware(handler)

	arg0 := apps.Arg0(args, apps.RELAY)
	return wtf.Serve(arg0, handler)
}

func Handler(dir string) http.Handler {
	return &webdav.Handler{
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WebDAV error: %s, %s\n", r.Method, err)
			}
		},
	}
}
