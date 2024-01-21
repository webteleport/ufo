package ser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps/ser/handler"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	arg0 := Arg0(args, "https://ufo.k0s.io")
	h := handler.Handler(".")
	if arg0 == "local" {
		port := utils.EnvPort(":8000")
		log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
		return http.ListenAndServe(port, h)
	}
	return wtf.Serve(arg0, h)
}
