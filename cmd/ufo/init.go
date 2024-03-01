package main

import (
	"net"
	"os"
	"path/filepath"
	"strings"
)

func isValidDomainName(domain string) bool {
	_, err := net.LookupHost(domain)
	return err == nil
}

func setRelayEnv() {
	if os.Getenv("RELAY") != "" {
		return
	}
	exe := os.Args[0]
	exe = strings.ToLower(exe)
	exe = strings.TrimSuffix(exe, ".exe")
	// println("exe="+exe)
	name := filepath.Base(exe)
	if isValidDomainName(name) {
		// println("RELAY=" + name)
		os.Setenv("RELAY", name)
	}
}

func init() {
	setRelayEnv()
}
