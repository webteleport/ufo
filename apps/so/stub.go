//go:build js || wasip1

package so

import "errors"

func Run(args []string) error {
	return errors.New("so: not supported on this platform")
}
