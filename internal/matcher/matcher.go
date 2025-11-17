package matcher

import (
	"fmt"
	"path/filepath"
	"strings"
)

// MatchAccount returns the account that applies to the given directory.
func MatchAccount(absDirPath string, accounts map[string][]string, defaultAccount string) (string, error) {
	cleanDirPath := filepath.Clean(absDirPath)

	for account, paths := range accounts {
		for _, path := range paths {
			norm := normalise(path)
			if strings.HasPrefix(cleanDirPath, norm) {
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
