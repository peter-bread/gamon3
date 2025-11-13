package locator

import (
	"os"
	"path/filepath"
)

func LocalConfigPath() (string, error) {
	start, _ := os.Getwd()
	stop, _ := os.UserHomeDir()
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
	// TODO: Custom error?
	return "", os.ErrNotExist
}
