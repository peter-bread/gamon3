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

func LocalConfigPath(os LocalOS) (string, error) {
	start, err := os.Getwd()
	if err != nil {
		return "", err
	}

	stop, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	candidates := []string{".gamon3.yaml", ".gamon3.yml"}

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
