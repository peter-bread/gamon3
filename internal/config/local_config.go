package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

// LocalConfig is a structure that stores data from a local configuration file.
type LocalConfig struct {
	Account string `yaml:"account"`
}

// LoadLocalConfig attempts to read data from a YAML file at `path` and returns
// a new `LocalConfig` structure.
func LoadLocalConfig(path string) (*LocalConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data), yaml.DisallowUnknownField())
	var cfg LocalConfig
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("%s", yaml.FormatError(err, true, true))
	}
	return &cfg, nil
}
