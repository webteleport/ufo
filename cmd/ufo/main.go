package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btwiuse/multicall"
	"github.com/btwiuse/portmux"
	"github.com/btwiuse/pub"

	// "github.com/webteleport/ufo/apps/caddy"
	"github.com/webteleport/ufo/apps/agent"
	"github.com/webteleport/ufo/apps/basicauth"
	"github.com/webteleport/ufo/apps/cookies"
	"github.com/webteleport/ufo/apps/dl"
	"github.com/webteleport/ufo/apps/echo"
	"github.com/webteleport/ufo/apps/freeport"
	"github.com/webteleport/ufo/apps/gitd"
	"github.com/webteleport/ufo/apps/gopilot"
	"github.com/webteleport/ufo/apps/gotip"
	"github.com/webteleport/ufo/apps/grafana"
	"github.com/webteleport/ufo/apps/hdr"
	"github.com/webteleport/ufo/apps/hello"
	"github.com/webteleport/ufo/apps/intercept"
	"github.com/webteleport/ufo/apps/ip"
	"github.com/webteleport/ufo/apps/logbody"
	"github.com/webteleport/ufo/apps/login"
	"github.com/webteleport/ufo/apps/mini"
	"github.com/webteleport/ufo/apps/mmdb"
	"github.com/webteleport/ufo/apps/multi"
	"github.com/webteleport/ufo/apps/naive"
	"github.com/webteleport/ufo/apps/nc"
	"github.com/webteleport/ufo/apps/notfound"
	"github.com/webteleport/ufo/apps/pkgpath"
	"github.com/webteleport/ufo/apps/pocket"
	"github.com/webteleport/ufo/apps/proxy"
	"github.com/webteleport/ufo/apps/raw"
	"github.com/webteleport/ufo/apps/relay"
	"github.com/webteleport/ufo/apps/resolve"
	"github.com/webteleport/ufo/apps/ser"
	"github.com/webteleport/ufo/apps/so"
	"github.com/webteleport/ufo/apps/sowc"
	"github.com/webteleport/ufo/apps/sowcmux"
	"github.com/webteleport/ufo/apps/sows"
	"github.com/webteleport/ufo/apps/sowsmux"
	"github.com/webteleport/ufo/apps/sse"
	"github.com/webteleport/ufo/apps/teleport"
	"github.com/webteleport/ufo/apps/term"
	"github.com/webteleport/ufo/apps/upgrade"
	"github.com/webteleport/ufo/apps/upload"
	"github.com/webteleport/ufo/apps/v2ray"
	"github.com/webteleport/ufo/apps/version"
	"github.com/webteleport/ufo/apps/vsc"
	"github.com/webteleport/ufo/apps/who"
	"github.com/webteleport/ufo/apps/whois"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	err := Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

var cmdRun multicall.RunnerFuncMap = map[string]multicall.RunnerFunc{
	// "caddy":        caddy.Run,
	// auth
	"dev/cookies":   cookies.Run,
	"dev/basicauth": basicauth.Run,
	"dev/login":     login.Run,
	"dev/who":       who.Run,
	// multilistener
	"dev/multi":    multi.Run,
	"dev/pkgpath":  pkgpath.Run,
	"dev/notfound": notfound.Run,
	"dev/sse":      sse.Run,
	"dev/hdr":      hdr.Run,
	"dev/proxy":    proxy.Run,
	"dev/nc":       nc.Run,
	"dev/logbody":  logbody.Run,
	// Proxy
	"so":      so.Run,
	"sowc":    sowc.Run,
	"sows":    sows.Run,
	"sowcmux": sowcmux.Run,
	"sowsmux": sowsmux.Run,
	"naive":   naive.Run,
	"v2ray":   v2ray.Run,

	// Cli
	// find freeport
	"freeport": freeport.Run,
	// go toolchain
	"gotip": gotip.Run,
	// resolve endpoint
	"resolve": resolve.Run,

	// Services

	// get client ip
	"ip": ip.Run,
	// upload file
	"upload": upload.Run,
	// share binary
	"dl": dl.Run,
	// pub assets
	"pub": pub.Run,
	// serve directory
	"ser": ser.Run,
	// grafana
	"grafana": grafana.Run,
	// teleport
	"teleport": teleport.Run,
	// intercept
	"intercept": intercept.Run,
	// ws + http mux
	"portmux": portmux.Run,
	// hello world
	"hello": hello.Run,
	// raw
	"raw": raw.Run,
	// echo request
	"echo": echo.Run,

	// Core

	// relay
	"hub":    relay.Run,
	"relay":  relay.Run,
	"pocket": pocket.Run,
	"mini":   mini.Run,
	// version info
	"version": version.Run,
	// binary upgrade
	"upgrade": upgrade.Run,

	// Apps

	// gopilot api
	"gopilot": gopilot.Run,
	// web terminal
	"term": term.Run,
	// vscode server
	"vsc": vsc.Run,
	// whois lookup
	"whois": whois.Run,
	// mmdb lookup
	"mmdb": mmdb.Run,
	// git daemon
	"gitd": gitd.Run,
	// agent management
	"agent": agent.Run,
}

func Run(args []string) error {
	err := cmdRun.Run(os.Args[1:])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
