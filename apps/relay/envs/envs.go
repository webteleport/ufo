package envs

import (
	"fmt"

	"github.com/webteleport/utils"
)

var (
	HOST       = utils.EnvHost("localhost")
	CERT       = utils.EnvCert("localhost.pem")
	KEY        = utils.EnvKey("localhost-key.pem")
	PORT       = utils.EnvPort(":3000")
	UDP_PORT   = utils.EnvUDPPort(PORT)
	ALT_SVC    = utils.EnvAltSvc(fmt.Sprintf(`webteleport="%s"`, UDP_PORT))
	HTTP_PORT  = utils.LookupEnvPort("HTTP_PORT")
	HTTPS_PORT = utils.LookupEnvPort("HTTPS_PORT")
)
