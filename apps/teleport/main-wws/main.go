package main

import (
	"github.com/vmware-labs/wasm-workers-server/kits/go/worker"
	"github.com/webteleport/ufo/apps/teleport/handler"
)

func main() {
	worker.Serve(handler.Handler("https://example.com"))
}
