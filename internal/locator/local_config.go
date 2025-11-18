package locator

import (
	"fmt"
	"os"
	"path/filepath"
)

func LocalConfigPath() (string, error) {
	start, err := os.Getwd()
	if err != nil {
		return "", err
	}

	stop, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := start
	candidates := []string{".gamon3.yaml", ".gamon3.yml"}

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

	return "", fmt.Errorf("could not find a local config file starting from %q", start)
}
