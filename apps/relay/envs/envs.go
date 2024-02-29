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
	HTTP_PORT  = utils.EnvHTTPPort(PORT)
	UDP_PORT   = utils.EnvUDPPort(PORT)
	HTTPS_PORT = utils.LookupEnvPort("HTTPS_PORT")
	ALT_SVC    = utils.EnvAltSvc(fmt.Sprintf(`webteleport="%s"`, UDP_PORT))
)
