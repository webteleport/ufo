module github.com/webteleport/ufo

go 1.23.2

// replace github.com/btwiuse/multicall => ../multicall
// replace github.com/webteleport/utils => ../utils
// replace github.com/webteleport/wtf => ../wtf
// replace github.com/webteleport/webteleport => ../webteleport
// replace github.com/webteleport/relay => ../relay

require (
	connectrpc.com/connect v1.16.1
	github.com/btwiuse/better v0.0.0
	github.com/btwiuse/bingo v0.0.3
	github.com/btwiuse/dl v0.0.1
	github.com/btwiuse/gost v0.0.4
	github.com/btwiuse/multicall v0.0.5
	github.com/btwiuse/portmux v0.1.0
	github.com/btwiuse/pretty v0.2.1
	github.com/btwiuse/pub v0.3.9
	github.com/btwiuse/rng v0.0.1
	github.com/btwiuse/sse v0.0.1
	github.com/btwiuse/tags v0.0.2
	github.com/btwiuse/version v0.0.1
	github.com/caddyserver/certmagic v0.20.0
	github.com/creativeprojects/go-selfupdate v1.2.0
	github.com/fermyon/spin/sdk/go/v2 v2.2.0
	github.com/gin-gonic/gin v1.9.1
	github.com/hashicorp/yamux v0.1.1
	github.com/jpillora/go-echo-server v0.5.0
	github.com/mdp/qrterminal/v3 v3.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/quic-go/quic-go v0.48.1
	github.com/tidwall/gjson v1.14.4
	github.com/v2fly/v2ray-core/v5 v5.15.1
	github.com/vmware-labs/wasm-workers-server v1.7.0
	// github.com/webteleport/caddy-webteleport v0.0.1
	github.com/webteleport/auth v0.0.9
	github.com/webteleport/relay v0.4.36
	github.com/webteleport/utils v0.2.16
	github.com/webteleport/webteleport v0.5.35
	github.com/webteleport/wtf v0.1.27
	golang.org/x/net v0.30.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	k0s.io/pkg/agent v0.1.15
	k0s.io/pkg/asciitransport v0.1.15
	k0s.io/pkg/dial v0.1.15
	k0s.io/pkg/wrap v0.1.15
)

require (
	code.gitea.io/sdk/gitea v0.17.1 // indirect
	filippo.io/edwards25519 v1.0.0-rc.1.0.20210721174708-390f27c3be20 // indirect
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	github.com/ActiveState/termtest/conpty v0.5.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20230301210247-d84479c117d7 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/btwiuse/connect v0.0.5 // indirect
	github.com/btwiuse/forward v0.0.0 // indirect
	github.com/btwiuse/muxr v0.0.1 // indirect
	github.com/btwiuse/wsconn v0.0.3 // indirect
	github.com/bufbuild/protocompile v0.8.0 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/creack/pty v1.1.21 // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/dgryski/go-metro v0.0.0-20211217172704-adc40b04c140 // indirect
	github.com/docker/docker v27.3.1+incompatible // indirect
	github.com/ebfe/bcrypt_pbkdf v0.0.0-20140212075826-3c8d2dcb253a // indirect
	github.com/ebi-yade/altsvc-go v0.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-fed/httpsig v1.1.0 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20220408131628-aac992c27451 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.19.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-github/v30 v30.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20230821062121-407c9e7a662f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/jhump/protoreflect v1.15.6 // indirect
	github.com/jpillora/ansi v0.0.0-20170202005112-f496b27cd669 // indirect
	github.com/jpillora/requestlog v0.0.0-20181015073026-df8817be5f82 // indirect
	github.com/jpillora/sizestr v0.0.0-20160130011556-e2ea2fa42fb9 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kataras/basicauth v0.0.3 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/klauspost/reedsolomon v1.11.7 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42 // indirect
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mholt/acmez v1.2.0 // indirect
	github.com/miekg/dns v1.1.62 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo/v2 v2.13.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pires/go-proxyproto v0.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/webtransport-go v0.8.1-0.20241018022711-4ac2c9250e66 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20220411075957-e3b120b3f5fb // indirect
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
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/v2fly/BrowserBridge v0.0.0-20210430233438-0570fc1d7d08 // indirect
	github.com/v2fly/VSign v0.0.0-20201108000810-e2adc24bf848 // indirect
	github.com/v2fly/ss-bloomring v0.0.0-20210312155135-28617310f63e // indirect
	github.com/xanzy/go-gitlab v0.100.0 // indirect
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	github.com/xtaci/smux v1.5.24 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	gitlab.com/yawning/edwards25519-extra.git v0.0.0-20211229043746-2f91fcc9fbdb // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20220204003609-77af0cba934d // indirect
	go.starlark.net v0.0.0-20230612165344-9532f5667272 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	go4.org/netipx v0.0.0-20230303233057-f1b76eb4bb35 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/oauth2 v0.18.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/term v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k0s.io v0.1.15 // indirect
	k8s.io/apimachinery v0.29.2 // indirect
	nhooyr.io/websocket v1.8.17 // indirect
	rsc.io/qr v0.2.0 // indirect
)
