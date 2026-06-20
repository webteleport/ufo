//go:build js || wasip1

package v2ray

import "errors"

func Run(args []string) error {
	return errors.New("v2ray: not supported on this platform")
}
