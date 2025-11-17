package resolve

import (
	"testing"

	"github.com/peter-bread/gamon3/internal/data"
)

func Test_resolveEnv(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		account string
		gh      GhHosts
		want    Result
		wantErr bool
	}{
		{
			name:    "error if account is empty",
			account: "",
			wantErr: true,
		},
		{
			name:    "error if account is invalid",
			account: "invalid",
			gh: mockGhHosts{
				users: []string{"alice", "bob"},
			},
			wantErr: true,
		},
		{
			name:    "use env var account if it is valid",
			account: "alice",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			want: Result{
				Current:     "alice",
				Account:     "alice",
				SourceKind:  Env,
				SourceValue: data.EnvVarAccount,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := resolveEnv(tt.account, tt.gh)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("resolveEnv() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("resolveEnv() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("resolveEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
