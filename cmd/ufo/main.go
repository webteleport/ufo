package main

import (
	"log"
	"os"

	"github.com/btwiuse/multicall"
	"github.com/webteleport/server"
	"github.com/webteleport/ufo/apps/caddy"
	"github.com/webteleport/ufo/apps/cookies"
	"github.com/webteleport/ufo/apps/echo"
	"github.com/webteleport/ufo/apps/gos"
	"github.com/webteleport/ufo/apps/hdr"
	"github.com/webteleport/ufo/apps/hello"
	"github.com/webteleport/ufo/apps/login"
	"github.com/webteleport/ufo/apps/metrics"
	"github.com/webteleport/ufo/apps/multi"
	"github.com/webteleport/ufo/apps/nc"
	"github.com/webteleport/ufo/apps/notfound"
	"github.com/webteleport/ufo/apps/pkgpath"
	"github.com/webteleport/ufo/apps/so"
	"github.com/webteleport/ufo/apps/sows"
	"github.com/webteleport/ufo/apps/sse"
	"github.com/webteleport/ufo/apps/teleport"
	"github.com/webteleport/ufo/apps/term"
	"github.com/webteleport/ufo/apps/who"

	_ "github.com/webteleport/utils/hack/quic-go-disable-receive-buffer-warning"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	err := Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

var cmdRun multicall.RunnerFuncMap = map[string]multicall.RunnerFunc{
	"cookies":      cookies.Run,
	"caddy":        caddy.Run,
	"pkgpath":      pkgpath.Run,
	"notfound":     notfound.Run,
	"hello":        hello.Run,
	"echo":         echo.Run,
	"login":        login.Run,
	"hdr":          hdr.Run,
	"nc":           nc.Run,
	"so":           so.Run,
	"sows":         sows.Run,
	"sse":          sse.Run,
	"term":         term.Run,
	"who":          who.Run,
	"metrics":      metrics.Run,
	"multi":        multi.Run,
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
