package nc

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/btwiuse/wsconn"
	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Run(args []string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsconn.Hijack(w)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = io.Copy(conn, io.TeeReader(conn, os.Stderr))
		if err != nil {
			log.Println(err)
		}
	})
	return wtf.Serve(apps.Arg0(args, apps.RELAY), utils.GinLoggerMiddleware(http.DefaultServeMux))
}
