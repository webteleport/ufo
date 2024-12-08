package raw

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/webteleport"
)

func Run([]string) error {
	ln, err := webteleport.Listen(context.Background(), apps.RELAY)
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ln.Addr().String())
	for {
		conn, err := ln.Accept()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handle(conn)
	}
	return nil
}

func handle(conn io.ReadWriteCloser) {
	defer conn.Close()
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Println("read request error:", err)
		return
	}
	log.Println(req)
	// body := "HTTP/1.1 200 OK\r\n\r\nHello, World!"
	body := "HTTP/1.1 204 No Content\r\n\r\n"
	conn.Write([]byte(body))
}
