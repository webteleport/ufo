// this version of pocket doesn't support connect proxy
package pocket

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/webteleport/relay"
	"github.com/webteleport/utils"
)

var (
	HOST = utils.EnvHost("^localhost$")
)

func Run(args []string) error {
	os.Args = append([]string{"pocket"}, args...)

	app := pocketbase.New()

	var hooksDir string
	app.RootCmd.PersistentFlags().StringVar(
		&hooksDir,
		"hooksDir",
		"",
		"the directory with the JS app hooks",
	)

	app.RootCmd.ParseFlags(args)

	// load jsvm (pb_hooks and pb_migrations)
	jsvm.MustRegister(app, jsvm.Config{
		HooksDir: hooksDir,
	})

	// registers the relay middleware
	app.OnServe().Bind(&hook.Handler[*core.ServeEvent]{
		Func: func(se *core.ServeEvent) error {
			log.Println("starting the relay server", "HOST", HOST)

			store := relay.NewSessionStore()

			if os.Getenv("LOGGIN") != "" {
				store.Use(utils.GinLoggerMiddleware)
			}

			mini := relay.NewWSServer(HOST, store)

			se.Router.BindFunc(func(re *core.RequestEvent) error {
				isPocketbaseHost := mini.IsRootExternal(re.Event.Request)
				isPocketbaseAPI := strings.HasPrefix(re.Event.Request.URL.Path, "/api/")
				isPocketbaseUI := strings.HasPrefix(re.Event.Request.URL.Path, "/_/")
				isPocketbase := isPocketbaseHost && (isPocketbaseAPI || isPocketbaseUI)

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
	})

	return app.Start()
}
