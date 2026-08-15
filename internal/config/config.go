package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/atomicswe/atproxy/internal/request"
)

type Config struct {
	Validator request.ValidatorConfig `json:"validator"`
}

const (
	configFilePath = "config.json"
)

func LoadConfig() (*Config, error) {
	config, err := loadConfig()
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			return nil, err
		}
		return createNewConfig()
	}

	return config, nil
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func createNewConfig() (*Config, error) {
	config := Config{}

	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
