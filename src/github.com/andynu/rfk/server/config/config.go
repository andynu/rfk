package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/andynu/rfk/server/observer"
)

type ConfigType struct {
	DataPath                   string
	WeatherUndergroundKey      string
	WeatherUndergroundLocation string
}

var Config ConfigType

func init() {
	Config = ConfigType{}
	Config.DataPath = dataPath()
}

func Load(configPath *string) {
	if configPath == nil || *configPath == "" {
		defaultPath := "./config.json"
		configPath = &defaultPath
	}
	config, err := loadJsonConfig(*configPath)
	if err != nil {
		panic(err)
	}
	Config.DataPath = config["data_dir"]
	Config.WeatherUndergroundKey = config["weather_underground_key"]
	Config.WeatherUndergroundLocation = config["weather_underground_location"]

	observer.Notify("config.loaded", struct{}{})
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
