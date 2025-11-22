// Package runtime provides concrete implementations of the interfaces defined
// in the parent package.
package runtime

import "github.com/peter-bread/gamon3/internal/authswitch"

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
