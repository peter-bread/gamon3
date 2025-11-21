package runtime

import "github.com/peter-bread/gamon3/internal/executor"

type Executor struct {
	runner executor.Runner
}

func NewExecutor() *Executor {
	return &Executor{
		runner: Runner{},
	}
}

func (e *Executor) Switch(account string) error {
	return executor.Switch(e.runner, account)
}

func (e *Executor) SwitchIfNeeded(account, current string) error {
	if account != current {
		return e.Switch(account)
	}
	return nil
}
