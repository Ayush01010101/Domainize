package utlities

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pterm/pterm"
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

	configpath, err := ConfigPath()
	if err != nil {
		panic(err)
	}

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

	// Stream the download to disk. When the server reports a size we drive a
	// real progress bar from it; otherwise fall back to an indeterminate spinner.
	if resp.ContentLength > 0 {
		bar, _ := pterm.DefaultProgressbar.
			WithTotal(int(resp.ContentLength)).
			WithTitle("Downloading mkcert").
			Start()

		if _, err := io.Copy(out, io.TeeReader(resp.Body, &progressWriter{bar: bar})); err != nil {
			bar.Stop()
			panic(err)
		}
		bar.Stop()
	} else {
		spinner, _ := pterm.DefaultSpinner.Start("Downloading mkcert")
		if _, err := io.Copy(out, resp.Body); err != nil {
			spinner.Fail("Download failed")
			panic(err)
		}
		spinner.Success("Downloaded mkcert")
	}

	if err := out.Close(); err != nil {
		panic(err)
	}

	if err := os.Chmod(mkcertPath, 0755); err != nil {
		panic(err)
	}

	if output, err := runMkcert(filepath.Dir(configpath), mkcertPath, "-install"); err != nil {
		panic("failed to install mkcert local CA: " + string(output))
	}
}

// progressWriter advances a pterm progress bar by the number of bytes streamed
// through it. It discards the data itself, so it's meant to sit behind an
// io.TeeReader that already writes the real bytes to disk.
type progressWriter struct {
	bar *pterm.ProgressbarPrinter
}

func (w *progressWriter) Write(p []byte) (int, error) {
	w.bar.Add(len(p))
	return len(p), nil
}
