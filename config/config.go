package config

import (
	"encoding/json"
	"os"
)

// Load configuration file and parse the json content
func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}

// Config struct for use in the applications with some general values
type Config struct {
	Contact   string         `json:"contact"`
	Resolver1 string         `json:"resolver1"`
	Resolver2 string         `json:"resolver2"`
	Domains   []ConfigDomain `json:"domains"`
	Output    string         `json:"output"`
}

// ConfigDomain struct for domains to monitor.
type ConfigDomain struct {
	Name string `json:"name"`
}
