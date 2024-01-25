package gotip

import (
	"os"

	"github.com/btwiuse/dl/version"
)

func Run(args []string) error {
	os.Args = append([]string{"gotip"}, args...)
	version.RunTip()
	return nil
}
