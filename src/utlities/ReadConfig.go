package utlities

import (
	"encoding/json"
	"os"
)

// ReadConfig loads and parses the config file from disk.
func ReadConfig() (Config, error) {
	var config Config

	configPath, err := ConfigPath()
	if err != nil {
		return config, err
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(contents, &config); err != nil {
		return config, err
	}

	return config, nil
}
