package utlities

import (
	"os"
	"path/filepath"
)

func SetupConfig() bool {

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	configpath := filepath.Join(configDir, "domainize", "config.json")

	return true
}
