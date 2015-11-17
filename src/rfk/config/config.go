package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type ConfigType struct {
	DataPath string
}

var Config ConfigType

func init() {
	Config = ConfigType{}
	Config.DataPath = dataPath()
}

func dataPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	panicOnErr(err)
	host, err := os.Hostname()
	panicOnErr(err)
	return fmt.Sprintf("%s/data/%s", dir, host)
}

func panicOnErr(err error) {
	if err != nil {
		panic(fmt.Errorf("library: %v", err))
	}
}
