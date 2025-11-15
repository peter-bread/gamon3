package locator_test

import (
	"testing"

	"github.com/peter-bread/gamon3/internal/locator"
)

func TestEnvAccount(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		want  string
		want2 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := locator.EnvAccount()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("EnvAccount() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("EnvAccount() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
