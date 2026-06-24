package utlities

import (
	"os"
	"os/user"
	"path/filepath"
)

func ConfigPath() (string, error) {
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" && sudoUser != "root" {
		userInfo, err := user.Lookup(sudoUser)
		if err == nil && userInfo.HomeDir != "" {
			return filepath.Join(userInfo.HomeDir, ".config", "domainize", "config.json"), nil
		}
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "domainize", "config.json"), nil
}
