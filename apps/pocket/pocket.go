// this version of pocket doesn't support connect proxy
package pocket

import (
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
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
	app.OnServe().Bind(RelayHook)

	return app.Start()
}
