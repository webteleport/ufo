package main

import (
	"github.com/vmware-labs/wasm-workers-server/kits/go/worker"
	"github.com/webteleport/ufo/apps/basicauth/handler"
)

func main() {
	worker.Serve(handler.Handler())
}
