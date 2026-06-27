package utlities

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

// CheckDomainConfig checks whether the config already has a domain configured.
//
// If no domain exists yet (or there is no config file), it returns true with a
// nil slice so the link command can continue normally. If one or more domains
// already exist, it shows a yes/no interactive select (radio buttons) asking the
// user whether to overwrite. Choosing "Yes" clears the existing domains so the
// newly provided domain replaces them and returns true along with the list of
// removed domains (so the caller can also drop them from the hosts file);
// choosing "No" returns false so the caller can stop early without running the
// remaining functions.
func CheckDomainConfig() (bool, []string) {
	configPath, err := ConfigPath()
	if err != nil {
		panic(err)
	}

	contents, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		// No config file yet, nothing to overwrite.
		return true, nil
	} else if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(contents, &config); err != nil {
		panic(err)
	}

	// No domain configured yet, continue normally.
	if len(config.Domain) == 0 {
		return true, nil
	}

	existing := make([]string, 0, len(config.Domain))
	for domain := range config.Domain {
		existing = append(existing, domain)
	}

	pterm.Warning.Printfln("A domain is already configured: %s", strings.Join(existing, ", "))

	selected, err := pterm.DefaultInteractiveSelect.
		WithDefaultText("Overwrite the existing domain?").
		WithOptions([]string{"Yes", "No"}).
		WithDefaultOption("No").
		Show()
	if err != nil {
		panic(err)
	}

	if selected != "Yes" {
		pterm.Info.Println("Keeping the existing configuration.")
		return false, nil
	}

	// Overwrite: clear the existing domains so the new one replaces them.
	config.Domain = make(map[string]DomainConfig)

	updatedContents, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(configPath, updatedContents, 0o644); err != nil {
		panic(err)
	}

	return true, existing
}
