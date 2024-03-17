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
func getLatestVersionInfo(quality string) (*VersionInfo, error) {
	link := fmt.Sprintf("%s/api/latest/server-%s-%s-web/%s", VSCODE_UPDATE_URL, getOS(), getARCH(), quality)
	log.Println(link)
	resp, err := http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
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

// download the latest version of vscode server by quality
// /commit:be210b3a60c7d60030c1d3d92da00d008edf6ab9/server-linux-arm64-web/insider
func downloadVersion(quality string, commit string) (string, error) {
	link := fmt.Sprintf("%s/commit:%s/server-%s-%s-web/%s", VSCODE_UPDATE_URL, commit, getOS(), getARCH(), quality)
	log.Println(link)
	resp, err := http.Get(link)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	pattern := fmt.Sprintf("vscode-server-archive-%s-%s-%s-%s-*", commit, getOS(), getARCH(), quality)
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
