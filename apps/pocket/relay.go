package pocket

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/webteleport/relay"
	"github.com/webteleport/utils"
)

// disable the relay server by setting HOST to an empty string
var HOST = utils.EnvHost("")

// execute as late as possible to make user provided hooks run early
var PRIORITY = StringToInt(os.Getenv("PRIORITY"), 99998)

var RelayHook = &hook.Handler[*core.ServeEvent]{
	Func: func(se *core.ServeEvent) error {
		if HOST == "" {
			return se.Next()
		}

		log.Println("starting the relay server", "HOST", HOST)

		s := relay.DefaultWSServer(HOST)

		// s.RootHandler doesn't work with the gin logger middleware
		// it's recommended to use pb_hooks to log requests
		if os.Getenv("LOGGIN") != "" {
			s.Use(utils.GinLoggerMiddleware)
		}

		se.Router.BindFunc(func(re *core.RequestEvent) error {
			r := re.Event.Request
			isPocketbaseHost := s.IsRootExternal(r)
			isAPI := strings.HasPrefix(r.URL.Path, "/api/")
			isUI := strings.HasPrefix(r.URL.Path, "/_/")
			isPocketbase := isPocketbaseHost && (isAPI || isUI)

			if os.Getenv("DEBUG") != "" {
				log.Println(fmt.Sprintf("isPocketbase (%v) := isPocketbaseHost (%v) && (isAPI (%v) || isUI (%v))", isPocketbase, isPocketbaseHost, isAPI, isUI))
			}

			// route non pocketbase requests to relay
			if !isPocketbase {
				s.ServeHTTP(re.Event.Response, re.Event.Request)
				return nil
			}

			return re.Next()
		})

		return se.Next()
	},
	Priority: PRIORITY,
}

func StringToInt(str string, fallback int) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return fallback
	}
	return num
}
