package utlities

import (
	"os"
	"path/filepath"
)

func Isconfigpresent() bool {

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	mkcertDir := filepath.Join(configDir, "domainize", "bin")
	_, configpatherror := os.Stat(mkcertDir)

	if configpatherror != nil {
		return false
	}
	return true

}
