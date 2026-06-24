package functions

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
)

func UpdateConfig(port int, domain string) {
	configPath, err := utlities.ConfigPath()
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
		panic(err)
	}

	contents, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		contents = []byte(`{"name":"domainize"}`)
	} else if err != nil {
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
