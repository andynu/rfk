package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func Load(configPath string) {
	config, err := loadJsonConfig(configPath)
	if err != nil {
		panic(err)
	}
	Config.DataPath = config["data_dir"]
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

func loadJsonConfig(jsonConfigFile string) (map[string]string, error) {
	type jsonConfig map[string]string
	var config jsonConfig

	file, err := ioutil.ReadFile(jsonConfigFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
