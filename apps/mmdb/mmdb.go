package mmdb

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/webteleport/ufo/apps"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func extractLastSegment(s string) string {
	// Clean the path and extract the last segment
	lastSegment := path.Base(path.Clean(s))

	return lastSegment
}

func GetClientIP(r *http.Request) (clientIP string) {
	// Retrieve the client IP address from the request headers
	for _, x := range []string{
		r.Header.Get("X-Envoy-External-Address"),
		r.Header.Get("X-Real-IP"),
		r.Header.Get("X-Forwarded-For"),
		r.RemoteAddr,
	} {
		if x != "" {
			clientIP = x
			break
		}
	}
	return
}

func mmdbHandler(w http.ResponseWriter, r *http.Request) {
	// Get the path from the request URL
	path := r.URL.Path

	domain := extractLastSegment(path)

	if domain == "/" {
		domain = GetClientIP(r)
	}

	fmt.Println(domain)
	lookup := os.Getenv("MMDB_LOOKUP")
	if lookup == "" {
		lookup = "github.com/btwiuse/mmdb/lookup"
	}
	// Execute the lookup command
	cmd := exec.Command(lookup, domain)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		// Handle the error if the lookup command fails
		http.Error(w, "Failed to execute lookup command", http.StatusInternalServerError)
		return
	}

	// Set the response content type to plain text
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write the output of the whois command as the response
	fmt.Fprint(w, string(output))
}

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	var h http.Handler
	h = http.HandlerFunc(mmdbHandler)
	h = utils.GinLoggerMiddleware(h)
	return wtf.Serve(Arg0(args, apps.RELAY), h)
}
