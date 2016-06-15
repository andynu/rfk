package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/andynu/rfk/server/observer"
)

type ConfigType struct {
	WeatherUndergroundKey      string
	WeatherUndergroundLocation string
}

var Config ConfigType
var ConfigPath string
var DataPath string

func init() {
	Config = ConfigType{}
	scriptDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Errorf("No sciptDir! Crazy! %q", err))
	}

	dataPath := path.Join(scriptDir, "data")
	if !pathExists(dataPath) {
		wd, _ := os.Getwd()
		dataPath = path.Join(wd, "data")
	}
	DataPath = dataPath
}

func Load(configPath string) {
	ConfigPath = DefaultConfigPath(configPath)
	CreateDefaultConfig(ConfigPath)
	fmt.Println("Config: %q", ConfigPath)
	config, err := loadJsonConfig(ConfigPath)
	if err != nil {
		panic(err)
	}

	Config.WeatherUndergroundKey = config["weather_underground_key"]
	Config.WeatherUndergroundLocation = config["weather_underground_location"]

	observer.Notify("config.loaded", struct{}{})
}

func DefaultConfigPath(configPath string) string {
	if configPath == "" {
		configPath = path.Join(scriptDir(), "config.json")
	}
	if !pathExists(configPath) {
		wd, _ := os.Getwd()
		configPath = path.Join(wd, "config.json")
	}
	return configPath
}

func CreateDefaultConfig(configPath string) {
	if !pathExists(configPath) {
		fmt.Println(configPath)
		configExamplePath := path.Join(scriptDir(), "config.json.example")
		cp(configExamplePath, configPath)
		log.Printf("Missing config: created default at %q", configPath)
	}
}

func DefaultDataPath(dataPath string) string {
	if dataPath == "" {
		dataPath = path.Join(scriptDir(), "data")
	}
	if !pathExists(dataPath) {
		wd, _ := os.Getwd()
		dataPath = path.Join(wd, "data")
	}
	DataPath = dataPath
	return dataPath
}

func CreateDataPath(dataPath string) {
	if !pathExists(dataPath) {
		log.Printf("Missing data dir: creating at %q", dataPath)
		os.Mkdir(dataPath, 0760)
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

func scriptDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return (err == nil)
}

func cp(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}

	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
