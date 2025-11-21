package authswitch_test

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/authswitch"
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
		runner  authswitch.Runner
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
			gotErr := authswitch.Switch(tt.runner, tt.account)
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

func TestSwitchIfNeeded(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		runner  authswitch.Runner
		account string
		current string
		wantErr bool
	}{
		{
			name:    "switch if account is different to current",
			runner:  mockRunner{},
			account: "alice",
			current: "bob",
		},
		{
			name:    "do not switch if account is the same as current",
			runner:  mockRunner{},
			account: "alice",
			current: "alice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := authswitch.SwitchIfNeeded(tt.runner, tt.account, tt.current)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SwitchIfNeeded() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SwitchIfNeeded() succeeded unexpectedly")
			}
		})
	}
}
