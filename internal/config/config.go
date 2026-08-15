package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/atomicswe/atproxy/internal/request"
	"github.com/atomicswe/atproxy/internal/server"
)

const (
	configFilePath = "config.json"
)

type Config struct {
	Server    server.ServerConfig     `json:"server"`
	Validator request.ValidatorConfig `json:"validator"`
}

// LoadConfig tries to load the config from the `configFilePath`
// if it fails due the file not existing, it creates a new config.
// At the end, it saves the config as it is loaded.
func LoadConfig() (*Config, error) {
	config := newConfig()
	err := load(config)
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return nil, err
	}

	return config, config.saveConfig()
}

func newConfig() *Config {
	return &Config{
		Server:    server.NewServerConfig(),
		Validator: request.NewValidatorConfig(),
	}
}

func load(c *Config) error {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) saveConfig() error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, data, 0644)

	return err
}
