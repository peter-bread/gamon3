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

// CurrentUser returns the current GH CLI user.
func (g *GhHosts) CurrentUser() string {
	return g.GitHubCom.User
}

// AllUsers returns a list of all authenticated GH CLI users.
func (g *GhHosts) AllUsers() []string {
	users := make([]string, 0, len(g.GitHubCom.Users))
	for user := range g.GitHubCom.Users {
		users = append(users, user)
	}
	return users
}
