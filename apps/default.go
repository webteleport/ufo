package apps

import "os"

var RELAY = EnvRELAY("https://ufo.k0s.io")

func EnvRELAY(s string) string {
	if relay := os.Getenv("RELAY"); relay != "" {
		return relay
	}
	return s
}

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Arg1(args []string, fallback string) string {
	if len(args) > 1 {
		return args[1]
	}
	return fallback
}
