//go:build js || wasip1

package mini

import "errors"

func Run(args []string) error {
	return errors.New("mini: not supported on this platform")
}
