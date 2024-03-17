package gotip

import (
	"os"
	"strings"

	"github.com/btwiuse/dl/version"
)

func Run(args []string) error {
	if gotip := os.Getenv("GOTIP"); gotip == "" {
		os.Args = append([]string{"gotip"}, args...)
		version.RunTip()
	} else {
		if !strings.HasPrefix(gotip, "go") {
			gotip = "go" + gotip
		}
		return RunVersion(gotip)(args)
	}
	return nil
}

func RunVersion(v string) func([]string) error {
	return func(args []string) error {
		os.Args = append([]string{v}, args...)
		version.Run(v)
		return nil
	}
}
