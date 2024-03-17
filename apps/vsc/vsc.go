package vsc

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/btwiuse/pretty"
	"github.com/phayes/freeport"
	"github.com/webteleport/ufo/apps/vsc/handler"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

type ServeWebArgs struct {
	Quality                  *string `json:"quality"`
	Host                     *string `json:"host"`
	SocketPath               *string `json:"socketPath"`
	Port                     *int    `json:"port"`
	ConnectionToken          *string `json:"connectionToken"`
	ConnectionTokenFile      *string `json:"connectionTokenFile"`
	WithoutConnectionToken   *bool   `json:"withoutConnectionToken"`
	AcceptServerLicenseTerms *bool   `json:"acceptServerLicenseTerms"`
	ServerBasePath           *string `json:"serverBasePath"`
	ServerDataDir            *string `json:"serverDataDir"`
	UserDataDir              *string `json:"userDataDir"`
	ExtensionsDir            *string `json:"extensionsDir"`
}

func Parse(args []string) (*ServeWebArgs, error) {
	flagSet := flag.NewFlagSet("serveWebArgs", flag.ContinueOnError)

	serveWebArgs := &ServeWebArgs{
		Quality:                  flagSet.String("quality", "insider", "Quality (stable | insider | exploration), defaults to 'insider'"),
		Host:                     flagSet.String("host", "127.0.0.1", "Host to listen on, defaults to '127.0.0.1'"),
		SocketPath:               flagSet.String("socket-path", "", "The path to a socket file for the server to listen to."),
		Port:                     flagSet.Int("port", 0, "Port to listen on, defaults to 0. If 0 is passed a random free port is picked."),
		ConnectionToken:          flagSet.String("connection-token", "", "A secret that must be included with all requests."),
		ConnectionTokenFile:      flagSet.String("connection-token-file", "", "A file containing a secret that must be included with all requests."),
		WithoutConnectionToken:   flagSet.Bool("without-connection-token", true, "Run without a connection token. Only use this if the connection is secured by other means."),
		AcceptServerLicenseTerms: flagSet.Bool("accept-server-license-terms", true, "If set, the user accepts the server license terms and the server will be started without a user prompt."),
		ServerBasePath:           flagSet.String("server-base-path", "", "Specifies the path under which the web UI and the code server is provided."),
		ServerDataDir:            flagSet.String("server-data-dir", "", "Specifies the directory that server data is kept in."),
		UserDataDir:              flagSet.String("user-data-dir", "", "Specifies the directory that user data is kept in. Can be used to open multiple distinct instances of Code."),
		ExtensionsDir:            flagSet.String("extensions-dir", "", "Set the root path for extensions."),
	}

	err := flagSet.Parse(args)
	if err != nil {
		return nil, err
	}

	if *serveWebArgs.Port == 0 {
		randport, err := freeport.GetFreePort()
		if err != nil {
			return nil, err
		}
		*serveWebArgs.Port = randport
	}

	return serveWebArgs, nil
}

func Run(args []string) error {
	serveWebArgs, err := Parse(args)
	if err != nil {
		return err
	}

	fmt.Println(pretty.JSONStringLine(serveWebArgs))

	info, err := getLatestVersionInfo(*serveWebArgs.Quality)
	if err != nil {
		return err
	}
	fmt.Println(pretty.JSONStringLine(info))

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	targetDir := fmt.Sprintf("%s/.vsc/cli/serve-web/%s", home, info.Version)
	_ = os.MkdirAll(targetDir, 0755)

	if !isInstalled(*serveWebArgs.Quality, targetDir) {
		err = downloadAndExtract(*serveWebArgs.Quality, info.Version, targetDir)
		if err != nil {
			return err
		}
	}

	err = startVersion(serveWebArgs, targetDir)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("http://%s:%d", *serveWebArgs.Host, *serveWebArgs.Port)
	return wtf.Serve("https://ufo.k0s.io", utils.GinLoggerMiddleware(handler.Handler(addr)))
}

func downloadAndExtract(quality string, commit string, targetDir string) error {
	archive, err := downloadVersion(quality, commit)
	if err != nil {
		return err
	}
	err = extractArchive(archive, targetDir)
	if err != nil {
		return err
	}
	return nil
}

func executableName(quality string) string {
	switch quality {
	case "insider":
		return "code-server-insiders"
	case "exploration":
		return "code-server-exploration"
	}
	return "code-server"
}

func isInstalled(quality, path string) bool {
	executable := filepath.Join(path, "bin", executableName(quality))
	_, err := os.Stat(executable)
	return err == nil
}

func extractArchive(archive string, targetDir string) error {
	cmd := fmt.Sprintf(`tar -xvf "%s" -C "%s" --strip 1`, archive, targetDir)
	return exec.Command("sh", "-c", cmd).Run()
}

func startVersion(args *ServeWebArgs, path string) error {
	log.Printf("Starting server at %s\n", *args.Host)

	executable := filepath.Join(path, "bin", executableName(*args.Quality))
	log.Printf("Executable: %s\n", executable)

	cmd := exec.Command(executable, "--host", *args.Host, "--port", fmt.Sprint(*args.Port), "--accept-server-license-terms", fmt.Sprint(*args.AcceptServerLicenseTerms))

	// Set the input/output options of the command
	cmd.Stdin = nil
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if *args.ServerBasePath != "" {
		cmd.Args = append(cmd.Args, "--server-base-path", *args.ServerBasePath)
	}
	if *args.ServerDataDir != "" {
		cmd.Args = append(cmd.Args, "--server-data-dir", *args.ServerDataDir)
	}
	if *args.UserDataDir != "" {
		cmd.Args = append(cmd.Args, "--user-data-dir", *args.UserDataDir)
	}
	if *args.ExtensionsDir != "" {
		cmd.Args = append(cmd.Args, "--extensions-dir", *args.ExtensionsDir)
	}
	if *args.WithoutConnectionToken {
		cmd.Args = append(cmd.Args, "--without-connection-token")
	}
	if *args.ConnectionTokenFile != "" {
		cmd.Args = append(cmd.Args, "--connection-token-file", *args.ConnectionTokenFile)
	}

	log.Println(cmd.Args)

	// Start the command
	return cmd.Start()
}
