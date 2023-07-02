package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	repoPath = "."
	port     = ":8080"
)

func main() {
	http.HandleFunc("/git-receive-pack", handleGitReceivePack)
	http.HandleFunc("/git-upload-pack", handleGitUploadPack)

	log.Printf("Git server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleGitReceivePack(w http.ResponseWriter, r *http.Request) {
	handleGitRequest(w, r, "receive-pack")
}

func handleGitUploadPack(w http.ResponseWriter, r *http.Request) {
	handleGitRequest(w, r, "upload-pack")
}

func handleGitRequest(w http.ResponseWriter, r *http.Request, command string) {
	repo := filepath.Join(repoPath, r.URL.Path)

	cmd := exec.Command("git", command, "--stateless-rpc", repo)
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to execute Git command: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
