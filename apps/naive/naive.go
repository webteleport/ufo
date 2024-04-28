package naive

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/webteleport/relay"
	"github.com/webteleport/webteleport"
)

// local listener works
//
//	https_proxy=http://a:b@localhost:8080 curl https://google.com -v
//
// while the remote listener doesn't
//
//	http_proxy=https://naive.remotehost:pass@remotehost:443 curl http://google.com -v
func Run([]string) error {
	ln1, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", ":8080")

	ln2, err := webteleport.Listen(context.Background(), "https://ufo.k0s.io/naive?naive=1")
	if err != nil {
		return err
	}
	log.Println("🛸 listening on", fmt.Sprintf("%s://%s", ln2.Addr().Network(), ln2.Addr().String()))

	os.Setenv("CONNECT_VERBOSE", "1")
	h := relay.NewConnectHandler()

	go http.Serve(ln1, h)
	return http.Serve(ln2, h)
}
