package whois

import (
	"fmt"
	"net/http"
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

func whoisHandler(w http.ResponseWriter, r *http.Request) {
	// Get the path from the request URL
	path := r.URL.Path

	domain := extractLastSegment(path)

	// Execute the whois command
	cmd := exec.Command("whois", domain)
	output, err := cmd.Output()
	if err != nil {
		// Handle the error if the whois command fails
		http.Error(w, "Failed to execute whois command", http.StatusInternalServerError)
		return
	}

	// Set the response content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Write the output of the whois command as the response
	fmt.Fprint(w, string(output))
}

func Run(args []string) error {
	var h http.Handler
	h = http.HandlerFunc(whoisHandler)
	h = utils.GinLoggerMiddleware(h)
	return wtf.Serve(apps.Arg0(args, apps.RELAY), h)
}
