package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

// GhHosts is a structure that represents a subset of a GH CLI `hosts.yml`
// file.
type GhHosts struct {
	GitHubCom struct {
		Users map[string]any `yaml:"users"`
		User  string         `yaml:"user"`
	} `yaml:"github.com"`
}

// LoadGhHosts attempts to read data from a YAML file at `path` and returns
// a new `GhHosts` structure.
func LoadGhHosts(path string) (*GhHosts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var hosts GhHosts
	return &hosts, yaml.Unmarshal(data, &hosts)
}

// GetCurrentUser returns the current GH CLI user.
func (g *GhHosts) GetCurrentUser() string {
	return g.GitHubCom.User
}

// GetAllUsers returns a list of all authenticated GH CLI users.
func (g *GhHosts) GetAllUsers() []string {
	users := make([]string, 0, len(g.GitHubCom.Users))
	for user := range g.GitHubCom.Users {
		users = append(users, user)
	}
	return users
}
