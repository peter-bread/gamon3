package authswitch

type Switcher interface {
	Switch(account string) error
	SwitchIfNeeded(account, current string) error
}

type Runner interface {
	Run(name string, args ...string) (stderr string, err error)
}
