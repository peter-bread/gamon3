// Package matcher provides simple directory-to-account matching based on
// configured path prefixes. It is used to determine which logical account
// a given filesystem path belongs to, falling back to a default account
// when no match is found.
package matcher

import (
	"fmt"
	"path/filepath"
	"strings"
)

// MatchAccount returns the account associated with the given absolute
// directory path. It checks the provided accounts map, where each account
// key is mapped to one or more path prefixes. The first matching prefix
// determines the account. If no prefixes match, the defaultAccount is
// returned. An error is returned if an account key is empty or if the
// default account is not provided.
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
