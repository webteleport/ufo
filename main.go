package main

import (
	"log"
	"os"

	"github.com/btwiuse/multicall"
	"github.com/webteleport/ufo/apps/echo"
	"github.com/webteleport/ufo/apps/gos"
	"github.com/webteleport/ufo/apps/hdr"
	"github.com/webteleport/ufo/apps/hello"
	"github.com/webteleport/ufo/apps/nc"
	"github.com/webteleport/ufo/apps/sse"
	"github.com/webteleport/ufo/apps/teleport"
	"github.com/webteleport/webteleport/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	err := Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

var cmdRun multicall.RunnerFuncMap = map[string]multicall.RunnerFunc{
	"hello":        hello.Run,
	"echo":         echo.Run,
	"hdr":          hdr.Run,
	"nc":           nc.Run,
	"sse":          sse.Run,
	"fileserver":   gos.Run,      // renamed from "gos" to "fileserver"
	"gos":          gos.Run,      // TODO delete this
	"teleport":     teleport.Run, // renamed from "reverseproxy" to "teleport"
	"reverseproxy": teleport.Run, // TODO delete this
	"station":      server.Run,   // renamed from "server" to "station"
	"server":       server.Run,   // TODO delete this
}

func Run(args []string) error {
	return cmdRun.Run(os.Args[1:])
}
