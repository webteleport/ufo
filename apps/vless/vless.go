package vless

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/btwiuse/rng"
	"github.com/mdp/qrterminal/v3"
	"github.com/phayes/freeport"
	core "github.com/v2fly/v2ray-core/v5"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport"
)

type Inbound struct {
	Port           int             `json:"port"`
	Listen         string          `json:"listen"`
	Protocol       string          `json:"protocol"`
	Settings       InboundSettings `json:"settings"`
	StreamSettings StreamSettings  `json:"streamSettings"`
}

type InboundSettings struct {
	Decryption string   `json:"decryption"`
	Clients    []Client `json:"clients"`
}

type Client struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
}

type StreamSettings struct {
	Network string `json:"network"`
}

type Outbound struct {
	Protocol string                 `json:"protocol"`
	Settings map[string]interface{} `json:"settings"`
}

type Config struct {
	Inbounds  []Inbound  `json:"inbounds"`
	Outbounds []Outbound `json:"outbounds"`
}

type Builder interface {
	AddInbound(Inbound) Builder
	AddOutbound(Outbound) Builder
	Build() (Config, error)
}

type configBuilder struct {
	config Config
}

func NewConfigBuilder() Builder {
	return &configBuilder{}
}

func (b *configBuilder) AddInbound(inbound Inbound) Builder {
	b.config.Inbounds = append(b.config.Inbounds, inbound)
	return b
}

func (b *configBuilder) AddOutbound(outbound Outbound) Builder {
	b.config.Outbounds = append(b.config.Outbounds, outbound)
	return b
}

func (b *configBuilder) Build() (Config, error) {
	return b.config, nil
}

func BuildConfigJSON(port int, clients ...string) string {
	builder := NewConfigBuilder()

	// Build clients slice
	var clientList []Client
	for _, client := range clients {
		clientList = append(clientList, Client{ID: client, Level: 0})
	}

	// Create a single inbound with multiple clients
	inbound := Inbound{
		Port:     port,
		Listen:   "127.0.0.1",
		Protocol: "vless",
		Settings: InboundSettings{
			Decryption: "none",
			Clients:    clientList,
		},
		StreamSettings: StreamSettings{Network: "ws"},
	}

	builder.AddInbound(inbound)

	outbound := Outbound{
		Protocol: "freedom",
		Settings: make(map[string]interface{}),
	}

	builder.AddOutbound(outbound)
	config, err := builder.Build()

	if err != nil {
		fmt.Println("Error building config: ", err)
		return ""
	}

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling config to JSON: ", err)
		return ""
	}

	return string(jsonData)
}

func GenerateVlessURL(baseURL, userID string) (string, error) {
	// Parse the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %v", err)
	}

	// Extract the host and scheme from the base URL
	host := u.Host
	path := u.Path

	// Build the VLESS URL
	vlessURL := fmt.Sprintf("vless://%s@%s:443?encryption=none&over=tls&security=tls&sni=%s&fp=randomized&type=ws&host=%s&path=%s#%s", userID, host, host, host, path, host)

	return vlessURL, nil
}

func Run(args []string) error {
	randport, err := freeport.GetFreePort()
	if err != nil {
		return err
	}

	uid := rng.NewUUID()
	if id := os.Getenv("VLESS_UUID"); id != "" {
		uid = id
	}

	relayURL := "https://ufo.k0s.io"
	if relay := os.Getenv("RELAY"); relay != "" {
		relayURL = relay
	}

	configJSON := BuildConfigJSON(randport, uid)
	// fmt.Println(configJSON)
	config, err := core.LoadConfig(core.FormatJSON, strings.NewReader(configJSON))
	if err != nil {
		return err
	}
	// fmt.Println(config)

	server, err := core.New(config)
	if err != nil {
		return err
	}

	if err := server.Start(); err != nil {
		return err
	}
	defer server.Close()

	// Explicitly triggering GC to remove garbage from config loading.
	runtime.GC()

	ln, err := webteleport.Listen(context.Background(), relayURL)
	if err != nil {
		return err
	}
	vlessURL, err := GenerateVlessURL(webteleport.AsciiURL(ln), uid)
	if err != nil {
		return err
	}
	qrterminal.Generate(vlessURL, qrterminal.L, os.Stdout)
	fmt.Println(vlessURL)

	err = http.Serve(ln, utils.GinLoggerMiddleware(utils.ReverseProxy(fmt.Sprintf(":%d", randport))))
	if err != nil {
		return err
	}

	return nil
}