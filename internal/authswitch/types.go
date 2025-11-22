package authswitch

// Switcher is the interface that wraps methods that attempt to switch the active
// GitHub CLI account.
type Switcher interface {
	Switch(account string) error
	SwitchIfNeeded(account, current string) error
}

// Runner is the interface that wraps the Run method, which should be used
// to run external programs.
type Runner interface {
	Run(name string, args ...string) (stderr string, err error)
}
