package matcher

import (
	"fmt"
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
func MatchAccount(dir string, accounts map[string][]string, defaultAccount string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	for account, paths := range accounts {
		for _, path := range paths {
			norm := normalise(path)
			if strings.HasPrefix(abs, norm) {
				if account == "" {
					return "", fmt.Errorf("account key cannot be empty for path %q", norm)
				}
				return account, nil
			}
		}
	}

	if defaultAccount == "" {
		return "", fmt.Errorf("field 'default' cannot be empty")
	}

	return defaultAccount, nil
}
