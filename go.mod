module github.com/webteleport/ufo

go 1.22.0

// replace github.com/webteleport/webteleport => ../webteleport
// replace github.com/webteleport/wtf => ../wtf
// replace github.com/webteleport/relay => ../relay
// replace github.com/webteleport/utils => ../utils

require (
	connectrpc.com/connect v1.14.0
	github.com/btwiuse/bingo v0.0.3
	github.com/btwiuse/gost v0.0.4
	github.com/btwiuse/multicall v0.0.4
	github.com/btwiuse/portmux v0.1.0
	github.com/btwiuse/pretty v0.2.1
	github.com/btwiuse/pub v0.2.1
	github.com/btwiuse/sse v0.0.1
	github.com/caddyserver/certmagic v0.20.0
	github.com/fermyon/spin/sdk/go/v2 v2.1.0
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.6.0
	github.com/hashicorp/yamux v0.1.1
	github.com/jpillora/go-echo-server v0.5.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/quic-go/quic-go v0.41.0
	github.com/tidwall/gjson v1.17.0
	github.com/vmware-labs/wasm-workers-server v1.7.0
	// github.com/webteleport/caddy-webteleport v0.0.1
	github.com/webteleport/auth v0.0.6
	github.com/webteleport/relay v0.2.17
	github.com/webteleport/utils v0.2.7
	github.com/webteleport/webteleport v0.4.2
	github.com/webteleport/wtf v0.1.4
	golang.org/x/net v0.20.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.32.0
	k0s.io v0.1.14
	k0s.io/pkg/agent v0.1.14
	k0s.io/pkg/asciitransport v0.1.14
	k0s.io/pkg/dial v0.1.14
	k0s.io/pkg/exporter v0.1.14
	k0s.io/pkg/wrap v0.1.14
	nhooyr.io/websocket v1.8.10
)

require (
	github.com/btwiuse/dl v0.0.0-20240214180850-acd4fdda14de
	github.com/btwiuse/rng v0.0.0
	github.com/btwiuse/version v0.0.0
	github.com/creativeprojects/go-selfupdate v1.1.3
)

require (
	code.gitea.io/sdk/gitea v0.17.1 // indirect
	filippo.io/edwards25519 v1.0.0-rc.1.0.20210721174708-390f27c3be20 // indirect
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	github.com/ActiveState/termtest/conpty v0.5.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alecthomas/kingpin/v2 v2.4.0 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20230301210247-d84479c117d7 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/beevik/ntp v1.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btwiuse/better v0.0.0 // indirect
	github.com/btwiuse/tags v0.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/creack/pty v1.1.21 // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/dennwc/btrfs v0.0.0-20230312211831-a1f570bd01a1 // indirect
	github.com/dennwc/ioctl v1.0.0 // indirect
	github.com/digitalocean/godo v1.41.0 // indirect
	github.com/docker/docker v25.0.1+incompatible // indirect
	github.com/ebi-yade/altsvc-go v0.1.1 // indirect
	github.com/elazarl/goproxy v0.0.0-20231117061959-7cc037d33fb5 // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20231117061959-7cc037d33fb5 // indirect
	github.com/ema/qdisc v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-fed/httpsig v1.1.0 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20220408131628-aac992c27451 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.2.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-github/v30 v30.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20230821062121-407c9e7a662f // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-envparse v0.1.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hodgesds/perf-utils v0.7.0 // indirect
	github.com/illumos/go-kstat v0.0.0-20210513183136-173c9b0a9973 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/jpillora/ansi v1.0.2 // indirect
	github.com/jpillora/requestlog v1.0.0 // indirect
	github.com/jpillora/sizestr v1.0.0 // indirect
	github.com/jsimonetti/rtnetlink v1.3.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kataras/basicauth v0.0.3 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/libdns/digitalocean v0.0.0-20230728223659-4f9064657aea // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/lufia/iostat v1.2.1 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-xmlrpc v0.0.3 // indirect
	github.com/mdlayher/ethtool v0.1.0 // indirect
	github.com/mdlayher/genetlink v1.3.2 // indirect
	github.com/mdlayher/netlink v1.7.2 // indirect
	github.com/mdlayher/socket v0.4.1 // indirect
	github.com/mdlayher/wifi v0.1.0 // indirect
	github.com/mholt/acmez v1.2.0 // indirect
	github.com/miekg/dns v1.1.58 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo/v2 v2.13.0 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus-community/go-runit v0.1.0 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/node_exporter v1.7.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/webtransport-go v0.6.0 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/safchain/ethtool v0.3.0 // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/templexxx/cpu v0.0.7 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/xanzy/go-gitlab v0.95.2 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	github.com/xtaci/smux v1.5.16 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	gitlab.com/yawning/edwards25519-extra.git v0.0.0-20211229043746-2f91fcc9fbdb // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20220204003609-77af0cba934d // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/mock v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/oauth2 v0.16.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v1.0.0 // indirect
	k0s.io/pkg/rng v0.1.14 // indirect
	k8s.io/apimachinery v0.29.1 // indirect
)
