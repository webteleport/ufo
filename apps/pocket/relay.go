package pocket

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/webteleport/relay"
	"github.com/webteleport/utils"
)

var HOST = utils.EnvHost("")

var RelayHook = &hook.Handler[*core.ServeEvent]{
	Func: func(se *core.ServeEvent) error {
		if HOST == "" {
			return se.Next()
		}

		log.Println("starting the relay server", "HOST", HOST)

		mini := relay.DefaultWSServer(HOST)

		if os.Getenv("LOGGIN") != "" {
			mini.Use(utils.GinLoggerMiddleware)
		}

		se.Router.BindFunc(func(re *core.RequestEvent) error {
			r := re.Event.Request
			isPocketbaseHost := mini.IsRootExternal(r)
			isAPI := strings.HasPrefix(r.URL.Path, "/api/")
			isUI := strings.HasPrefix(r.URL.Path, "/_/")
			isPocketbase := isPocketbaseHost && (isAPI || isUI)

			if os.Getenv("VERBOSE") != "" {
				log.Println(fmt.Sprintf("isPocketbase (%v) := isPocketbaseHost (%v) && (isAPI (%v) || isUI (%v))", isPocketbase, isPocketbaseHost, isAPI, isUI))
			}

			// route non pocketbase requests to relay
			if !isPocketbase {
				mini.ServeHTTP(re.Event.Response, re.Event.Request)
				return nil
			}

			return re.Next()
		})

		return se.Next()
	},
	Priority: -99999, // execute as early as possible
}
