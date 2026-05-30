//go:build js || wasip1

package sows

import "errors"

func Run(args []string) error {
	return errors.New("sows: not supported on this platform")
}
