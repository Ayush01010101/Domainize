package utlities

import (
	"os"
	"path/filepath"
)

func Isconfigpresent() bool {

	configpath, err := ConfigPath()
	if err != nil {
		panic(err)
	}

	mkcertDir := filepath.Join(filepath.Dir(configpath), "bin")
	_, configpatherror := os.Stat(mkcertDir)

	if configpatherror != nil {
		return false
	}
	return true

}
