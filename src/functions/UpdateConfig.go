package functions

import (
	"encoding/json"
	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
	"os"
	"path/filepath"
)

func UpdateConfig(port int, domain string) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(configDir, "domainize", "config.json")

	contents, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config utlities.Config
	if err := json.Unmarshal(contents, &config); err != nil {
		panic(err)
	}

	if config.Name == "" {
		config.Name = "domainize"
	}

	if config.Domain == nil {
		config.Domain = make(map[string]utlities.DomainConfig)
	}

	config.Domain[domain] = utlities.DomainConfig{
		HTTPS: true,
		Port:  port,
	}

	updatedContents, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(configPath, updatedContents, 0o644); err != nil {
		panic(err)
	}
}
