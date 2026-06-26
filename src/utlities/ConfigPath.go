package utlities

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func ConfigPath() (string, error) {
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" && sudoUser != "root" {
		userInfo, err := user.Lookup(sudoUser)
		if err == nil && userInfo.HomeDir != "" {
			var configDir string
			switch runtime.GOOS {
			case "darwin":
				configDir = filepath.Join(userInfo.HomeDir, "Library", "Application Support")
			case "windows":
				configDir = filepath.Join(userInfo.HomeDir, "AppData", "Roaming")
			default:
				configDir = filepath.Join(userInfo.HomeDir, ".config")
			}
			return filepath.Join(configDir, "domainize", "config.json"), nil
		}
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "domainize", "config.json"), nil
}
