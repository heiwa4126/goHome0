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

// Write down configuration in JSON format
func (conf *Config) Write() error {

	file, dir, err := GetPath()
	if err != nil {
		return err
	}

	err = CreateDirIfNotExist(dir)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(*conf, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, bytes, 0644)
}

// ReadJSONFile はJSONファイルを読み込む
func ReadJSONFile(filename string, out interface{}) error {

	reader, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	return json.NewDecoder(reader).Decode(out)
}

// Read ...
func Read() (*Config, error) {

	file, _, err := GetPath()
	if err != nil {
		return nil, err
	}

	var conf Config
	err = ReadJSONFile(file, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
