package matcher

import (
	"os"
	"path/filepath"
	"strings"
)

func normalise(path string) string {
	if path == "" {
		return path
	}

	if path == "~" {
		if home, err := os.UserHomeDir(); err == nil {
			return home
		}
	}

	if strings.HasPrefix(path, "~/") {
		if home, err := os.UserHomeDir(); err == nil {
			path = filepath.Join(home, path[2:])
		}
	}

	return filepath.Clean(os.ExpandEnv(path))
}
