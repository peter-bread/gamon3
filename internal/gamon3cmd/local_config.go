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
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/goccy/go-yaml"
)

type LocalConfig struct {
	Account string `yaml:"account"`
}

func (l *LocalConfig) Load(path string, allowedUsers []string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, l); err != nil {
		return err
	}

	if err := l.Validate(allowedUsers); err != nil {
		return err
	}

	return nil
}

func (l *LocalConfig) Validate(allowedUsers []string) error {
	if err := l.ValidateUsers(allowedUsers); err != nil {
		return err
	}

	return nil
}

func (l *LocalConfig) ValidateUsers(allowedUsers []string) error {
	if l.Account == "" {
		return fmt.Errorf("%s", "local config: 'account' field is required")
	}

	if !slices.Contains(allowedUsers, l.Account) {
		return fmt.Errorf("%s: '%s' %s", "local config", l.Account, "has not been registered with GH CLI")
	}

	return nil
}

func GetLocalConfigPath() (string, error) {
	start, _ := os.Getwd()
	stop, _ := os.UserHomeDir()
	dir := start
	candidates := []string{".gamon.yaml", ".gamon.yml"}
	for {

		for _, name := range candidates {
			candidate := filepath.Join(dir, name)
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
		}

		parent := filepath.Dir(dir)

		if dir == stop || parent == dir {
			return "", fmt.Errorf("%s", "Could not find a local config file")
		}

		dir = parent
	}
}
