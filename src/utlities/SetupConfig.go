package utlities

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Name   string                  `json:"name"`
	Domain map[string]DomainConfig `json:"domain,omitempty"`
}

type DomainConfig struct {
	HTTPS     bool             `json:"https"`
	Port      int              `json:"port"`
	Subdomain []map[int]string `json:"subdomain,omitempty"`
}

func SetupConfig() bool {

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	configpath := filepath.Join(configDir, "domainize", "config.json")

	data := Config{
		Name: "domainize",
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
