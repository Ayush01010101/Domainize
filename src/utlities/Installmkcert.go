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
	fmt.Println("installmkcert called")
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

	configpath, err := ConfigPath()
	if err != nil {
		panic(err)
	}
	fmt.Println("config directory", filepath.Dir(configpath))
	mkcertDir := filepath.Join(filepath.Dir(configpath), "bin")
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
