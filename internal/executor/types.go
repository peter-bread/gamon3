package executor

type Executor interface {
	Switch(account string) error
	SwitchIfNeeded(account, current string) error
}
