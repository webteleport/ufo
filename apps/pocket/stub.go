//go:build js || wasip1

package pocket

import "errors"

func Run(args []string) error {
	return errors.New("pocket: not supported on this platform")
}
