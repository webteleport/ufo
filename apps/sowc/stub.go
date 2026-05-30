//go:build js || wasip1

package sowc

import "errors"

func Run(args []string) error {
	return errors.New("sowc: not supported on this platform")
}
