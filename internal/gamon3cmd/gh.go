package gamon3cmd

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type GHHosts struct {
	GitHubCom struct {
		Users map[string]any `yaml:"users"`
		User  string         `yaml:"user"`
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

func (g *GHHosts) GetCurrentUser() string {
	return g.GitHubCom.User
}

func (g *GHHosts) GetAllUsers() []string {
	users := make([]string, 0, len(g.GitHubCom.Users))
	for user := range g.GitHubCom.Users {
		users = append(users, user)
	}
	return users
}

func GetGHHostsPath() (string, error) {
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
