package main

import (
	"github.com/vmware-labs/wasm-workers-server/kits/go/worker"
	"github.com/webteleport/ufo/apps/ip/handler"
)

func main() {
	worker.Serve(handler.Handler())
}
