package utlities

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SetupConfig() bool {

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	configpath := filepath.Join(configDir, "domainize", "config.json")

	// Some dummy config data to seed the file.
	data := map[string]any{
		"name":    "domainize",
		"version": "1.0.0",
		"domains": []string{},
	}

	contents, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(configpath, contents, 0o644); err != nil {
		panic(err)
	}

	return true
}
