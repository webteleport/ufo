package sowc

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"

	"github.com/btwiuse/wsdial"
	"github.com/webteleport/ufo/apps"
)

// CAVEAT: curl recognizes only lowercase http_proxy, not HTTP_PROXY
func proxyHints(addr string) {
	log.Println("# bash")
	println(fmt.Sprintf("export HTTP_PROXY=http://127.0.0.1%s HTTPS_PROXY=http://127.0.0.1%s", addr, addr))
	println(fmt.Sprintf("export http_proxy=http://127.0.0.1%s https_proxy=http://127.0.0.1%s", addr, addr))
	println(fmt.Sprintf("export HTTP_PROXY=socks5h://127.0.0.1%s HTTPS_PROXY=socks5h://127.0.0.1%s", addr, addr))
	println(fmt.Sprintf("export http_proxy=socks5h://127.0.0.1%s https_proxy=socks5h://127.0.0.1%s", addr, addr))

	log.Println("# cmd")
	println(fmt.Sprintf("set HTTP_PROXY=http://127.0.0.1%s HTTPS_PROXY=http://127.0.0.1%s", addr, addr))
	println(fmt.Sprintf("set http_proxy=http://127.0.0.1%s https_proxy=http://127.0.0.1%s", addr, addr))
	println(fmt.Sprintf("set HTTP_PROXY=socks5h://127.0.0.1%s HTTPS_PROXY=socks5h://127.0.0.1%s", addr, addr))
	println(fmt.Sprintf("set http_proxy=socks5h://127.0.0.1%s https_proxy=socks5h://127.0.0.1%s", addr, addr))

	log.Println("# powershell")
	println(fmt.Sprintf("$env:HTTP_PROXY='http://127.0.0.1%s'; $env:HTTPS_PROXY='http://127.0.0.1%s'", addr, addr))
	println(fmt.Sprintf("$env:http_proxy='http://127.0.0.1%s'; $env:https_proxy='http://127.0.0.1%s'", addr, addr))
	println(fmt.Sprintf("$env:HTTP_PROXY='socks5h://127.0.0.1%s'; $env:HTTPS_PROXY='socks5h://127.0.0.1%s'", addr, addr))
	println(fmt.Sprintf("$env:http_proxy='socks5h://127.0.0.1%s'; $env:https_proxy='socks5h://127.0.0.1%s'", addr, addr))
}

func Run(args []string) error {
	addr := apps.Arg0(args, ":8123")
	remote := apps.Arg1(args, "wss://sows.ufo.k0s.io")
	log.Println("[arg0] socks5 listening on", addr)
	log.Println("[arg1] remote socks5+wss", remote)
	proxyHints(addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	ep, _ := url.Parse(remote)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			wsconn, err := wsdial.Dial(ep)
			if err != nil {
				log.Println(err)
				return
			}
			go io.Copy(wsconn, conn)
			io.Copy(conn, wsconn)
		}()
	}
}
