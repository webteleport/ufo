package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btwiuse/multicall"

	// "github.com/webteleport/ufo/apps/caddy"
	"github.com/webteleport/ufo/apps/basicauth"
	"github.com/webteleport/ufo/apps/cookies"
	"github.com/webteleport/ufo/apps/dl"
	"github.com/webteleport/ufo/apps/echo"
	"github.com/webteleport/ufo/apps/freeport"
	"github.com/webteleport/ufo/apps/gopilot"
	"github.com/webteleport/ufo/apps/ser"
	"github.com/webteleport/ufo/apps/hdr"
	"github.com/webteleport/ufo/apps/hello"
	"github.com/webteleport/ufo/apps/login"
	"github.com/webteleport/ufo/apps/metrics"
	"github.com/webteleport/ufo/apps/multi"
	"github.com/webteleport/ufo/apps/nc"
	"github.com/webteleport/ufo/apps/notfound"
	"github.com/webteleport/ufo/apps/pkgpath"
	"github.com/webteleport/ufo/apps/proxy"
	"github.com/webteleport/ufo/apps/relay"
	"github.com/webteleport/ufo/apps/so"
	"github.com/webteleport/ufo/apps/sowc"
	"github.com/webteleport/ufo/apps/sowcmux"
	"github.com/webteleport/ufo/apps/sows"
	"github.com/webteleport/ufo/apps/sowsmux"
	"github.com/webteleport/ufo/apps/sse"
	"github.com/webteleport/ufo/apps/teleport"
	"github.com/webteleport/ufo/apps/term"
	"github.com/webteleport/ufo/apps/who"
	"github.com/webteleport/ufo/apps/whois"

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
	"cookies": cookies.Run,
	// "caddy":        caddy.Run,
	"pkgpath":      pkgpath.Run,
	"notfound":     notfound.Run,
	"dl":           dl.Run,
	"proxy":        proxy.Run,
	"hello":        hello.Run,
	"gopilot":      gopilot.Run,
	"basicauth":    basicauth.Run,
	"echo":         echo.Run,
	"login":        login.Run,
	"hdr":          hdr.Run,
	"hub":          relay.Run,
	"relay":        relay.Run,
	"nc":           nc.Run,
	"so":           so.Run,
	"freeport":     freeport.Run,
	"sowc":         sowc.Run,
	"sows":         sows.Run,
	"sowcmux":      sowcmux.Run,
	"sowsmux":      sowsmux.Run,
	"sse":          sse.Run,
	"term":         term.Run,
	"who":          who.Run,
	"whois":        whois.Run,
	"metrics":      metrics.Run,
	"multi":        multi.Run,
	"ser":          ser.Run,
	"teleport":     teleport.Run, // renamed from "reverseproxy" to "teleport"
	"reverseproxy": teleport.Run, // TODO delete this
}

func Run(args []string) error {
	err := cmdRun.Run(os.Args[1:])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
