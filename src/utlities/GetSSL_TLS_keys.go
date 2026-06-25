package utlities

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

func mkcertBinaryName() string {
	if runtime.GOOS == "windows" {
		return "mkcert.exe"
	}
	return "mkcert"
}

func GetSSL_TLS_keys(domain string) (string, string, error) {
	configPath, err := ConfigPath()
	if err != nil {
		return "", "", err
	}

	configDir := filepath.Dir(configPath)
	certFile := filepath.Join(configDir, domain+".pem")
	keyFile := filepath.Join(configDir, domain+"-key.pem")

	// Check if PEM files already exist
	if _, err := os.Stat(certFile); err == nil {
		if _, err := os.Stat(keyFile); err == nil {
			return certFile, keyFile, nil
		}
	}

	mkcertPath := filepath.Join(configDir, "bin", mkcertBinaryName())

	output, err := runMkcert(configDir, mkcertPath, domain)
	fmt.Print("output", string(output))
	if err != nil {
		return "", "", err
	}

	return certFile, keyFile, nil
}

func runMkcert(configDir string, mkcertPath string, args ...string) ([]byte, error) {
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" && sudoUser != "root" {
		userInfo, err := user.Lookup(sudoUser)
		if err == nil && userInfo.HomeDir != "" {
			cmdArgs := []string{
				"-u", sudoUser,
				"env",
				"HOME=" + userInfo.HomeDir,
				"XDG_CONFIG_HOME=" + filepath.Join(userInfo.HomeDir, ".config"),
				mkcertPath,
			}
			cmdArgs = append(cmdArgs, args...)

			cmd := exec.Command("sudo", cmdArgs...)
			cmd.Dir = configDir
			return cmd.CombinedOutput()
		}
	}

	cmd := exec.Command(mkcertPath, args...)
	cmd.Dir = configDir
	return cmd.CombinedOutput()
}
