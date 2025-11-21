package executor_test

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/executor"
)

type mockRunner struct {
	stderr string
	err    error
}

func (m mockRunner) Run(name string, args ...string) (string, error) {
	return m.stderr, m.err
}

func TestSwitch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		runner  executor.Runner
		account string
		wantErr bool
	}{
		{
			name: "error if run fails",
			runner: mockRunner{
				err: fmt.Errorf("run failed"),
			},
			account: "alice",
			wantErr: true,
		},
		{
			name:    "return nil if run succeeds",
			runner:  mockRunner{},
			account: "alice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := executor.Switch(tt.runner, tt.account)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Switch() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Switch() succeeded unexpectedly")
			}
		})
	}
}
