package utlities

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func Installmkcert() {
	var downloadURL string
	mkcertBinary := "mkcert"

	switch runtime.GOOS {
	case "linux":
		downloadURL = "https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-" + runtime.GOARCH

	case "windows":
		downloadURL = "https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-windows-" + runtime.GOARCH + ".exe"
		mkcertBinary = "mkcert.exe"

	case "darwin":
		downloadURL = "https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-darwin-" + runtime.GOARCH

	default:
		panic("unsupported operating system: " + runtime.GOOS)
	}

	// os.UserConfigDir() resolves the OS-appropriate config root:
	//   Linux:   $XDG_CONFIG_HOME or $HOME/.config
	//   macOS:   $HOME/Library/Application Support
	//   Windows: %AppData% (C:\Users\<user>\AppData\Roaming)

	configDir, err := os.UserConfigDir()
	fmt.Println("config directory", configDir)
	if err != nil {
		panic(err)
	}
	mkcertDir := filepath.Join(configDir, "domainize", "bin")
	mkcertPath := filepath.Join(mkcertDir, mkcertBinary)

	resp, err := http.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("failed to download mkcert: " + resp.Status)
	}

	if err := os.MkdirAll(mkcertDir, 0755); err != nil {
		panic(err)
	}

	out, err := os.Create(mkcertPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		panic(err)
	}

	os.Chmod(mkcertPath, 0755)
}
