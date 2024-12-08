package vsc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

type VersionInfo struct {
	Name           string `json:"name"`
	Version        string `json:"version"`
	ProductVersion string `json:"productVersion"`
	Timestamp      int64  `json:"timestamp"`
}

const VSCODE_UPDATE_URL = "https://update.code.visualstudio.com"

func getOS() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "win32"
	default:
		return os
	}
}

func getARCH() string {
	switch arch := runtime.GOARCH; arch {
	case "amd64":
		return "x64"
	default:
		return arch
	}
}

// quality: stable, insider
func (args *ServeWebArgs) getLatestVersionInfo() (*VersionInfo, error) {
	link := fmt.Sprintf("%s/api/latest/server-%s-%s-web/%s", VSCODE_UPDATE_URL, getOS(), getARCH(), *args.Quality)
	if *args.Verbose {
		log.Println(link)
	}
	resp, err := http.Get(link)
	if err != nil {
		fallbackVersion := &VersionInfo{Version: "latest"}
		log.Println(fmt.Errorf("failed to make request: %w, falling back to: %v", err, fallbackVersion))
		return fallbackVersion, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var versionInfo VersionInfo
	err = json.Unmarshal(body, &versionInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &versionInfo, nil
}

// download the latest version of vscode server by quality, example:
// /commit:be210b3a60c7d60030c1d3d92da00d008edf6ab9/server-linux-arm64-web/insider
func (args *ServeWebArgs) downloadVersion(commit string) (string, error) {
	link := fmt.Sprintf("%s/commit:%s/server-%s-%s-web/%s", VSCODE_UPDATE_URL, commit, getOS(), getARCH(), *args.Quality)
	if *args.Verbose {
		log.Println(link)
	}
	resp, err := http.Get(link)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	pattern := fmt.Sprintf("vscode-server-archive-%s-%s-%s-%s-*", commit, getOS(), getARCH(), *args.Quality)
	file, err := os.CreateTemp(os.TempDir(), pattern)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return file.Name(), nil
}
