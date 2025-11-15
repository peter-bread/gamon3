package core

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

// MatchAccount returns the account that applies to the given directory.
func MatchAccount(dir string, accounts map[string][]string, defaultAccount string) string {
	if dir == "" {
		return defaultAccount
	}

	abs, err := filepath.Abs(dir)
	if err != nil {
		// If something goes wrong resolving the path, fall back to default.
		return defaultAccount
	}

	for account, paths := range accounts {
		for _, path := range paths {
			norm := normalise(path)
			if strings.HasPrefix(abs, norm) {
				return account
			}
		}
	}

	return defaultAccount
}
