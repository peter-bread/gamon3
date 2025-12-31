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

// Package runtime provides concrete implementations of the interfaces defined
// in the parent package.
package runtime

import "github.com/peter-bread/gamon3/v2/internal/authswitch"

type Switcher struct {
	runner authswitch.Runner
}

// NewSwitcher returns a new instance of Switcher with a command Runner.
func NewSwitcher() *Switcher {
	return &Switcher{
		runner: Runner{},
	}
}

// Switch calls `gh auth switch --user <account>`.
func (s *Switcher) Switch(account string) error {
	return authswitch.Switch(s.runner, account)
}

// SwitchIfNeeded calls Switch only if the current account differs from the account that should
// be active.
func (s *Switcher) SwitchIfNeeded(account, current string) error {
	return authswitch.SwitchIfNeeded(s.runner, account, current)
}
