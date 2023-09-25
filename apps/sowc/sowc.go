package sowc

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"

	"k0s.io/pkg/dial"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Arg1(args []string, fallback string) string {
	if len(args) > 1 {
		return args[1]
	}
	return fallback
}

func Run(args []string) error {
	addr := Arg0(args, ":8123")
	remote := Arg1(args, "wss://sows.ufo.k0s.io")
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			wsconn, err := dial.Dial(ep)
			if err != nil {
				log.Println(err)
				return
			}
			go io.Copy(wsconn, conn)
			io.Copy(conn, wsconn)
		}()
	}
}