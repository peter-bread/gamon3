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

package resolve

import (
	"fmt"

	"github.com/peter-bread/gamon3/internal/matcher"
)

func resolveMain(path string, gh GhHosts, loader MainConfigLoader, os OS) (Result, error) {
	cfg, err := loader.Load(path)
	if err != nil {
		return Result{}, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return Result{}, err
	}

	// This is the account that should be used based on the current directory.
	account, err := matcher.MatchAccount(cwd, cfg.Accounts, cfg.Default)
	if err != nil {
		return Result{}, fmt.Errorf("main config %s", err)
	}

	if !isValidGitHubAccount(account, gh) {
		return Result{}, fmt.Errorf("main config account %q is not authenticated", account)
	}

	return Result{
		Current:     gh.CurrentUser(),
		Account:     account,
		SourceKind:  Main,
		SourceValue: path,
	}, nil
}
