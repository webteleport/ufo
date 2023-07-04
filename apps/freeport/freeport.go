package freeport

import (
	"fmt"

	"github.com/phayes/freeport"
)

func Run(args []string) error {
	port, err := freeport.GetFreePort()
	if err != nil {
		return err
	}
	fmt.Println(port)
	return nil
	// port is ready to listen on
}
