package vsc

import (
	"flag"
	"fmt"
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
	Relay                    *string `json:"relay"`
	Log                      *string `json:"log"`
	Verbose                  *bool   `json:"verbose"`
	Version                  *bool   `json:"version"`
	Quality                  *string `json:"quality"`
	Host                     *string `json:"host"`
	SocketPath               *string `json:"socketPath"`
	Port                     *int    `json:"port"`
	ConnectionToken          *string `json:"connectionToken"`
	ConnectionTokenFile      *string `json:"connectionTokenFile"`
	WithoutConnectionToken   *bool   `json:"withoutConnectionToken"`
	AcceptServerLicenseTerms *bool   `json:"acceptServerLicenseTerms"` // ignored
	ServerBasePath           *string `json:"serverBasePath"`
	ServerDataDir            *string `json:"serverDataDir"`
	UserDataDir              *string `json:"userDataDir"`
	ExtensionsDir            *string `json:"extensionsDir"`
}

func Parse(args []string) (*ServeWebArgs, error) {
	flagSet := flag.NewFlagSet("serveWebArgs", flag.ContinueOnError)

	serveWebArgs := &ServeWebArgs{
		Relay:                  flagSet.String("relay", "https://ufo.k0s.io", "Relay URL."),
		Log:                    flagSet.String("log", "off", "Log level: {off,critical,error,warn,info,debug,trace}, defaults to 'off'."),
		Verbose:                flagSet.Bool("verbose", false, "Verbose logging."),
		Version:                flagSet.Bool("v", false, "Show version."),
		Quality:                flagSet.String("quality", "insider", "Quality: {insider,stable,exploration}, defaults to 'insider'"),
		Host:                   flagSet.String("host", "127.0.0.1", "Host to listen on, defaults to '127.0.0.1'"),
		SocketPath:             flagSet.String("socket-path", "", "The path to a socket file for the server to listen to."),
		Port:                   flagSet.Int("port", 0, "Port to listen on, defaults to 0. If 0 is passed a random free port is picked."),
		ConnectionToken:        flagSet.String("connection-token", "", "A secret that must be included with all requests."),
		ConnectionTokenFile:    flagSet.String("connection-token-file", "", "A file containing a secret that must be included with all requests."),
		WithoutConnectionToken: flagSet.Bool("without-connection-token", false, "Run without a connection token. Only use this if the connection is secured by other means."),
		ServerBasePath:         flagSet.String("server-base-path", "", "Specifies the path under which the web UI and the code server is provided."),
		ServerDataDir:          flagSet.String("server-data-dir", "", "Specifies the directory that server data is kept in."),
		UserDataDir:            flagSet.String("user-data-dir", "", "Specifies the directory that user data is kept in. Can be used to open multiple distinct instances of Code."),
		ExtensionsDir:          flagSet.String("extensions-dir", "", "Set the root path for extensions."),
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

	if *serveWebArgs.Verbose {
		fmt.Println(pretty.JSONStringLine(serveWebArgs))
	}

	info, err := serveWebArgs.getLatestVersionInfo()
	if err != nil {
		return err
	}
	if *serveWebArgs.Verbose {
		fmt.Println(pretty.JSONStringLine(info))
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	targetDir := fmt.Sprintf("%s/.vsc/cli/serve-web/%s", home, info.Version)
	_ = os.MkdirAll(targetDir, 0755)

	if !serveWebArgs.isInstalled(targetDir) {
		archive, err := serveWebArgs.downloadVersion(info.Version)
		if err != nil {
			return err
		}
		err = extractArchive(archive, targetDir)
		if err != nil {
			return err
		}
	}

	if *serveWebArgs.Version {
		serveWebArgs.showVersion(targetDir)
		os.Exit(0)
	}

	go serveWebArgs.startVersion(targetDir)

	addr := fmt.Sprintf("http://%s:%d", *serveWebArgs.Host, *serveWebArgs.Port)
	return wtf.Serve(*serveWebArgs.Relay, utils.GinLoggerMiddleware(handler.Handler(addr)))
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

func (args *ServeWebArgs) isInstalled(path string) bool {
	executable := filepath.Join(path, "bin", executableName(*args.Quality))
	_, err := os.Stat(executable)
	return err == nil
}

// TODO support windows/zip
func extractArchive(archive string, targetDir string) error {
	cmd := fmt.Sprintf(`tar -xvf "%s" -C "%s" --strip 1`, archive, targetDir)
	return exec.Command("sh", "-c", cmd).Run()
}

func (args *ServeWebArgs) startVersion(path string) error {
	executable := filepath.Join(path, "bin", executableName(*args.Quality))

	cmd := exec.Command(executable,
		"--host", *args.Host,
		"--port", fmt.Sprint(*args.Port),
		"--log", fmt.Sprint(*args.Log),
		"--accept-server-license-terms",
	)

	if *args.Version {
		cmd.Args = append(cmd.Args, "-v")
	}

	// Set the input/output options of the command
	cmd.Stdin = nil
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if *args.SocketPath != "" {
		cmd.Args = append(cmd.Args, "--socket-path", *args.SocketPath)
	}
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
	if *args.ConnectionToken != "" {
		cmd.Args = append(cmd.Args, "--connection-token", *args.ConnectionToken)
	}
	if *args.ConnectionTokenFile != "" {
		cmd.Args = append(cmd.Args, "--connection-token-file", *args.ConnectionTokenFile)
	}

	if *args.Verbose {
		fmt.Println(cmd.String())
	}

	// Start the command
	return cmd.Start()
}

func (args *ServeWebArgs) showVersion(path string) error {
	executable := filepath.Join(path, "bin", executableName(*args.Quality))

	cmd := exec.Command(executable, "-v")

	// Set the input/output options of the command
	cmd.Stdin = nil
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Start the command
	return cmd.Run()
}
