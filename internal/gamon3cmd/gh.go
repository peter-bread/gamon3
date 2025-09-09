/*
Copyright © 2025 Peter Sheehan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

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
