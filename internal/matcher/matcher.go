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
