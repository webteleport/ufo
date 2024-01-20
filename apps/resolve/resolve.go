package resolve

import (
	"log"
	"net/url"

	"github.com/webteleport/webteleport"
)

func resolve(s string) {
	log.Println(s)
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	log.Println("ENV:", webteleport.ENV("ALT_SVC"))
	log.Println("TXT:", webteleport.TXT(u.Host))
	log.Println("HEAD:", webteleport.HEAD(u.String()))
	es := webteleport.Resolve(u)
	log.Println("endpoints:", es)
	log.Println()
}

func Run(args []string) error {
	resolve("https://ufo.k0s.io")
	resolve("https://k0s.io")
	resolve("https://hk.k0s.io")
	resolve("https://k1s.io")
	for _, a := range args {
		resolve(a)
	}
	return nil
}
