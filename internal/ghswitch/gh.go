package ghswitch

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type GHHosts struct {
	GitHubCom struct {
		User string `yaml:"user"`
	} `yaml:"github.com"`
}

func (g *GHHosts) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, g); err != nil {
		return err
	}
	return nil
}

func GetGHConfigPath() (string, error) {
	if configDir, found := os.LookupEnv("GH_CONFIG_DIR"); found {
		return filepath.Join(configDir, "hosts.yml"), nil
	}

	if xdgConfigDir, found := os.LookupEnv("XDG_CONFIG_HOME"); found {
		return filepath.Join(xdgConfigDir, "gh", "hosts.yml"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "gh", "hosts.yml"), nil
}
