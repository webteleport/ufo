package resolve

import (
	"log"
	"net/url"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport/endpoint"
)

func resolve(s string) {
	s = utils.AsURL(s)
	log.Println(s)
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	log.Println("ENV:", endpoint.AltSvcFromEnv("ALT_SVC"))
	log.Println("HEAD:", endpoint.AltSvcFromHEAD(u.String()))

	log.Println("endpoint.Resolve", u)
	es := endpoint.Resolve(u)
	log.Println("endpoints:", es)
	log.Println()
}

func Run(args []string) error {
	resolve(apps.RELAY)
	resolve("https://k0s.io")
	resolve("https://hk.k0s.io")
	resolve("https://k1s.io")
	for _, a := range args {
		resolve(a)
	}
	return nil
}
