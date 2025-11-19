package runtime

import "github.com/peter-bread/gamon3/internal/executor"

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Switch(account string) error {
	return executor.Switch(account)
}

func (e *Executor) SwitchIfNeeded(account, current string) error {
	if account != current {
		return e.Switch(account)
	}
	return nil
}
