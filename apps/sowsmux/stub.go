//go:build js || wasip1

package sowsmux

import "errors"

func Run(args []string) error {
	return errors.New("sowsmux: not supported on this platform")
}
