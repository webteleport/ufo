//go:build js || wasip1

package wagi

import "errors"

func Run(args []string) error {
	return errors.New("wagi: not supported on this platform")
}
