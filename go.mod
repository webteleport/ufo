module github.com/webteleport/ufo

go 1.24

// replace github.com/btwiuse/multicall => ../multicall
// replace github.com/webteleport/utils => ../utils
// replace github.com/webteleport/wtf => ../wtf
// replace github.com/webteleport/webteleport => ../webteleport
// replace github.com/webteleport/relay => ../relay

require (
	connectrpc.com/connect v1.18.1
	github.com/btwiuse/better v0.0.0
	github.com/btwiuse/bingo v0.0.5 // indirect
	github.com/btwiuse/dl v0.0.1
	github.com/btwiuse/gost v0.0.4
	github.com/btwiuse/multicall v0.0.5
	github.com/btwiuse/portmux v0.1.0
	github.com/btwiuse/pretty v0.2.1
	github.com/btwiuse/proxy v0.0.0
	github.com/btwiuse/pub v0.3.10
	github.com/btwiuse/rng v0.0.1
	github.com/btwiuse/sse v0.0.1
	github.com/btwiuse/tags v0.0.2
	github.com/btwiuse/version v0.0.1
	github.com/btwiuse/wsconn v0.0.3
	github.com/btwiuse/wsdial v0.0.0
	github.com/caddyserver/certmagic v0.20.0
	github.com/chadgpt/gopilot v0.0.4
	github.com/creativeprojects/go-selfupdate v1.2.0
	github.com/fermyon/spin/sdk/go/v2 v2.2.0
	github.com/hashicorp/yamux v0.1.2
	github.com/jpillora/go-echo-server v0.5.0
	github.com/mdp/qrterminal/v3 v3.2.0
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/pocketbase/pocketbase v0.25.4
	github.com/quic-go/quic-go v0.49.0
	github.com/vmware-labs/wasm-workers-server v1.7.0
	// github.com/webteleport/caddy-webteleport v0.0.1
	github.com/webteleport/auth v0.0.9
	github.com/webteleport/relay v0.4.52-alpha.1
	github.com/webteleport/utils v0.2.19-alpha.1
	github.com/webteleport/webteleport v0.5.40-alpha.1
	github.com/webteleport/wtf v0.1.31-alpha.1
	github.com/xtls/xray-core v1.8.24
	golang.org/x/net v0.35.0
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.5
	k0s.io/pkg/agent v0.1.18-alpha.8
	k0s.io/pkg/asciitransport v0.1.18-alpha.8
)

require (
	code.gitea.io/sdk/gitea v0.17.1 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	github.com/ActiveState/termtest/conpty v0.5.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20230301210247-d84479c117d7 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go-v2 v1.36.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.8 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.28.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.51 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.28 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.48 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.6.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.72.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.14 // indirect
	github.com/aws/smithy-go v1.22.2 // indirect
	github.com/btwiuse/connect v0.0.5 // indirect
	github.com/btwiuse/dispatcher v0.0.0 // indirect
	github.com/btwiuse/forward v0.0.0 // indirect
	github.com/btwiuse/muxr v0.0.1 // indirect
	github.com/cloudflare/circl v1.4.0 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/creack/pty v1.1.21 // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/denisbrodbeck/machineid v1.0.1 // indirect
	github.com/dgryski/go-metro v0.0.0-20211217172704-adc40b04c140 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.11.4 // indirect
	github.com/docker/docker v27.3.1+incompatible // indirect
	github.com/domodwyer/mailyak/v3 v3.6.2 // indirect
	github.com/dop251/base64dec v0.0.0-20231022112746-c6c9f9a96217 // indirect
	github.com/dop251/goja v0.0.0-20241009100908-5f46f2705ca3 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20240728170619-29b559befffc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ebi-yade/altsvc-go v0.1.1 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/ganigeorgiev/fexpr v0.4.1 // indirect
	github.com/ghodss/yaml v1.0.1-0.20220118164431-d8423dcdf344 // indirect
	github.com/go-fed/httpsig v1.1.0 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20220408131628-aac992c27451 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-github/v30 v30.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20241029153458-d1b30febd7db // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jpillora/ansi v0.0.0-20170202005112-f496b27cd669 // indirect
	github.com/jpillora/requestlog v0.0.0-20181015073026-df8817be5f82 // indirect
	github.com/jpillora/sizestr v0.0.0-20160130011556-e2ea2fa42fb9 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kataras/basicauth v0.0.3 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mholt/acmez v1.2.0 // indirect
	github.com/miekg/dns v1.1.63 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/onsi/ginkgo/v2 v2.21.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pires/go-proxyproto v0.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pocketbase/dbx v1.11.0 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/webtransport-go v0.8.1-0.20241018022711-4ac2c9250e66 // indirect
	github.com/refraction-networking/utls v1.6.7 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sagernet/sing v0.4.1 // indirect
	github.com/sagernet/sing-shadowsocks v0.2.7 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20240715131351-a2f2c23f1771 // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/templexxx/cpu v0.0.7 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/v2fly/ss-bloomring v0.0.0-20210312155135-28617310f63e // indirect
	github.com/vishvananda/netlink v1.3.0 // indirect
	github.com/vishvananda/netns v0.0.4 // indirect
	github.com/xanzy/go-gitlab v0.100.0 // indirect
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	github.com/xtaci/smux v1.5.16 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	github.com/xtls/reality v0.0.0-20240712055506-48f0b2d5ed6d // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	gitlab.com/yawning/edwards25519-extra.git v0.0.0-20211229043746-2f91fcc9fbdb // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20220204003609-77af0cba934d // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	go4.org/netipx v0.0.0-20231129151722-fdeea329fbba // indirect
	gocloud.dev v0.40.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/exp v0.0.0-20250210185358-939b2ce775ac // indirect
	golang.org/x/image v0.24.0 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/oauth2 v0.26.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/term v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	golang.zx2c4.com/wintun v0.0.0-20230126152724-0fa3db229ce2 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20231211153847-12269c276173 // indirect
	google.golang.org/api v0.220.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gvisor.dev/gvisor v0.0.0-20231202080848-1f7806d17489 // indirect
	k0s.io v0.1.18-alpha.8 // indirect
	k8s.io/apimachinery v0.32.0 // indirect
	lukechampine.com/blake3 v1.3.0 // indirect
	modernc.org/libc v1.55.3 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.8.2 // indirect
	modernc.org/sqlite v1.34.5 // indirect
	nhooyr.io/websocket v1.8.17 // indirect
	rsc.io/qr v0.2.0 // indirect
)

tool github.com/btwiuse/bingo/cmd/bingo
