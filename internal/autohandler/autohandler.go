// Package autohandler is a minimal drop-in replacement for the small
// slice of btwiuse/gost that ufo's so/sows/sowsmux apps actually use:
// a single Handle(conn) method that peeks the first byte of the
// connection and dispatches to a SOCKS4, SOCKS5, or HTTP CONNECT
// handler. It depends only on github.com/go-gost/gosocks4 and
// github.com/go-gost/gosocks5, both of which are already pulled in as
// transitive dependencies.
package autohandler

import (
	"bufio"
	"io"
	"net"
	"net/http"

	"github.com/go-gost/gosocks4"
	"github.com/go-gost/gosocks5"
	socks5srv "github.com/go-gost/gosocks5/server"
)

// Handler dispatches a raw net.Conn to the appropriate protocol
// handler based on the first byte the peer sends.
type Handler struct{}

func New() *Handler { return &Handler{} }

// Handle implements the auto-dispatch behaviour previously provided by
// btwiuse/gost.AutoHandler().Handle(conn).
func (h *Handler) Handle(conn net.Conn) error {
	br := bufio.NewReader(conn)
	b, err := br.Peek(1)
	if err != nil {
		return err
	}
	// Wrap the buffered reader so protocol handlers see a Conn that
	// reads from the buffered side and writes straight to the socket,
	// matching the behaviour btwiuse/gost's autoHandler relies on.
	wrapped := &bufferedConn{Conn: conn, r: br}

	switch b[0] {
	case gosocks4.Ver4:
		return handleSOCKS4(wrapped)
	case gosocks5.Ver5:
		return socks5srv.DefaultHandler.Handle(wrapped)
	default:
		return handleHTTP(wrapped)
	}
}

func handleSOCKS4(conn net.Conn) error {
	req, err := gosocks4.ReadRequest(conn)
	if err != nil {
		return err
	}
	if req.Cmd != gosocks4.CmdConnect {
		return gosocks4.NewReply(gosocks4.Rejected, nil).Write(conn)
	}
	upstream, err := net.Dial("tcp", req.Addr.String())
	if err != nil {
		return gosocks4.NewReply(gosocks4.Failed, nil).Write(conn)
	}
	defer upstream.Close()
	if err := gosocks4.NewReply(gosocks4.Granted, nil).Write(conn); err != nil {
		return err
	}
	return pipe(conn, upstream)
}

func handleHTTP(conn net.Conn) error {
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if req.Method != http.MethodConnect {
		resp := &http.Response{
			StatusCode: http.StatusMethodNotAllowed,
			ProtoMajor: 1,
			ProtoMinor: 1,
		}
		return resp.Write(conn)
	}

	upstream, err := net.Dial("tcp", req.Host)
	if err != nil {
		resp := &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			ProtoMajor: 1,
			ProtoMinor: 1,
		}
		return resp.Write(conn)
	}
	defer upstream.Close()

	if _, err := io.WriteString(conn, "HTTP/1.1 200 Connection established\r\n\r\n"); err != nil {
		return err
	}
	return pipe(conn, upstream)
}

func pipe(a, b net.Conn) error {
	errCh := make(chan error, 2)
	go func() { _, err := io.Copy(b, a); errCh <- err }()
	go func() { _, err := io.Copy(a, b); errCh <- err }()
	<-errCh
	return nil
}

type bufferedConn struct {
	net.Conn
	r *bufio.Reader
}

func (c *bufferedConn) Read(b []byte) (int, error) { return c.r.Read(b) }
