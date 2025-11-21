package runtime

import "github.com/peter-bread/gamon3/internal/authswitch"

type Switcher struct {
	runner authswitch.Runner
}

func NewSwitcher() *Switcher {
	return &Switcher{
		runner: Runner{},
	}
}

func (s *Switcher) Switch(account string) error {
	return authswitch.Switch(s.runner, account)
}

func (s *Switcher) SwitchIfNeeded(account, current string) error {
	return authswitch.SwitchIfNeeded(s.runner, account, current)
}
