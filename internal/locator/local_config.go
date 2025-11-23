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

package locator

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalOS interface {
	Stat(name string) (os.FileInfo, error)
	Getwd() (dir string, err error)
	UserHomeDir() (string, error)
}

// LocalConfigPath searches upwards from the current directory, looking for a local config file.
// This file can be any one of: .gamon.yaml, .gamon.yml, .gamon3.yaml, or .gamon3.yml.
// In most cases, the search will stop after checking the user's home directory. If the directory
// where the search is starting is not a descendant of the home directory, then the search will go
// all the way to the filesystem root.
func LocalConfigPath(os LocalOS) (string, error) {
	start, err := os.Getwd()
	if err != nil {
		return "", err
	}

	stop, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	candidates := []string{".gamon.yaml", ".gamon.yml", ".gamon3.yaml", ".gamon3.yml"}

	return searchUpward(os, start, stop, candidates)
}

func searchUpward(os LocalOS, start string, stop string, candidates []string) (string, error) {
	dir := start

	for {
		for _, name := range candidates {
			candidate := filepath.Join(dir, name)
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
		}

		parent := filepath.Dir(dir)
		if dir == stop || parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not find any of %v starting from %q", candidates, start)
}
