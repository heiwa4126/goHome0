package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// AppName this application name
const AppName = "goHome0"

// Config ...
type Config struct {
	APIKey   string
	Endpoint string
}

// GetPath configファイルのパスを得る
func GetPath() (string, string, error) {

	var configFile string

	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}

	configDir := filepath.Join(home, ".config")
	configFile = filepath.Join(configDir, AppName+".json")

	return configFile, configDir, nil
}

// CreateDirIfNotExist ... (TODO: write comment here)
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// Write write down configuration in JSON format
func (conf *Config) Write() error {

	confFile, confDir, err := GetPath()
	if err != nil {
		return err
	}

	err = CreateDirIfNotExist(confDir)
	if err != nil {
		return err
	}

	j1, err := json.MarshalIndent(*conf, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(confFile, j1, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Read TODO: write comment here
func Read() (*Config, error) {

	var conf Config
	configFile, _, err := GetPath()
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(configFile)

	if err != nil {
		// configFileがない
		return nil, err
	}

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
