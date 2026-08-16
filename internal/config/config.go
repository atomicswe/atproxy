package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/atomicswe/atproxy/internal/request"
	"github.com/atomicswe/atproxy/internal/server"
)

const (
	configFileName = "atproxy.json"
)

type Config struct {
	Server    server.ServerConfig     `json:"server"`
	Validator request.ValidatorConfig `json:"validator"`
}

// LoadConfig tries to load the config from the config path
// if it fails due the file not existing, it creates a new config.
// At the end, it saves the config as it is loaded.
func LoadConfig() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	log.Println("config file location:", path)

	config := newConfig()
	err = load(path, config)
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return nil, err
	}

	return config, config.saveConfig(path)
}

func newConfig() *Config {
	return &Config{
		Server:    server.NewServerConfig(),
		Validator: request.NewValidatorConfig(),
	}
}

func load(path string, c *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) saveConfig(path string) error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)

	return err
}

func getConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir = filepath.Join(dir, "atproxy")
	if err = os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}
