package sowcmux

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"

	"github.com/hashicorp/yamux"
	"github.com/webteleport/ufo/apps"

	"github.com/btwiuse/wsdial"
)

func Run(args []string) error {
	addr := apps.Arg0(args, ":8123")
	remote := apps.Arg1(args, "wss://sowsmux.ufo.k0s.io")
	log.Println("socks5 listening on", addr)
	log.Println("remote socks5+wss", remote)
	log.Println("# bash")
	log.Println(fmt.Sprintf("export HTTP_PROXY=http://127.0.0.1%s HTTPS_PROXY=http://127.0.0.1%s", addr, addr))
	log.Println(fmt.Sprintf("export HTTP_PROXY=socks5h://127.0.0.1%s HTTPS_PROXY=socks5h://127.0.0.1%s", addr, addr))
	log.Println("# cmd")
	log.Println(fmt.Sprintf("set HTTP_PROXY=http://127.0.0.1%s HTTPS_PROXY=http://127.0.0.1%s", addr, addr))
	log.Println(fmt.Sprintf("set HTTP_PROXY=socks5h://127.0.0.1%s HTTPS_PROXY=socks5h://127.0.0.1%s", addr, addr))
	log.Println("# powershell")
	log.Println(fmt.Sprintf("$env:HTTP_PROXY='http://127.0.0.1%s'; $env:HTTPS_PROXY='http://127.0.0.1%s'", addr, addr))
	log.Println(fmt.Sprintf("$env:HTTP_PROXY='socks5h://127.0.0.1%s'; $env:HTTPS_PROXY='socks5h://127.0.0.1%s'", addr, addr))

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	ep, _ := url.Parse(remote)

	wsconn, err := wsdial.Dial(ep)
	if err != nil {
		return err
	}

	// Setup client side of yamux
	session, err := yamux.Client(wsconn, nil)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			stream, err := session.Open()
			if err != nil {
				log.Println(err)
				return
			}
			go io.Copy(stream, conn)
			io.Copy(conn, stream)
		}()
	}
}
