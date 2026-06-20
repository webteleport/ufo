//go:build js || wasip1

package relay

import "errors"

func Run(args []string) error {
	return errors.New("hub/relay: not supported on this platform")
}
