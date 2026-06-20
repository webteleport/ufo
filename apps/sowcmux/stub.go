//go:build js || wasip1

package sowcmux

import "errors"

func Run(args []string) error {
	return errors.New("sowcmux: not supported on this platform")
}
