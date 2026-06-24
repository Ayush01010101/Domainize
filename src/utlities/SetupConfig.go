package utlities

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("setup config called")
	configpath, err := ConfigPath()
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Dir(configpath), 0o755); err != nil {
		panic(err)
	}

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
