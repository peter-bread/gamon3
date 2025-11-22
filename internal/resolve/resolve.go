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

// Package resolve provides a mechanism for deciding which GitHub CLI account should
// be active.
package resolve

import (
	"fmt"
)

// Resolve decides which method should be used to determine current account and delegates to other functions.
// It returns information about the account, and any errors that occur.
func Resolve(l Locator, ghLoader GhHostsLoader, localLoader LocalConfigLoader, mainLoader MainConfigLoader, os OS) (Result, error) {
	ghHostsPath, err := l.GhHostsPath()
	if err != nil {
		return Result{}, err
	}

	ghHosts, err := ghLoader.Load(ghHostsPath)
	if err != nil {
		return Result{}, err
	}

	if envAccount, found := l.EnvAccount(); found {
		return resolveEnv(envAccount, ghHosts)
	}

	if localPath, err := l.LocalConfigPath(); err == nil {
		return resolveLocal(localPath, ghHosts, localLoader)
	}

	if mainPath, err := l.MainConfigPath(); err == nil {
		return resolveMain(mainPath, ghHosts, mainLoader, os)
	}

	return Result{}, fmt.Errorf("no config sources found to be able to resolve account")
}
