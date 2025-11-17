package resolve

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/config"
)

func Test_resolveLocal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		gh      GhHosts
		loader  LocalConfigLoader
		want    Result
		wantErr bool
	}{
		{
			name: "error if gh loader fails",
			loader: mockLocalLoader{
				err: fmt.Errorf("local config loader failed"),
			},
			wantErr: true,
		},
		{
			name: "error if account is empty",
			loader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "",
				},
			},
			wantErr: true,
		},
		{
			name: "error if account is invalid",
			loader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "invalid",
				},
			},
			gh: mockGhHosts{
				users: []string{"alice", "bob"},
			},
			wantErr: true,
		},
		{
			name: "good",
			path: "path/to/local/config",
			loader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "bob",
				},
			},
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			want: Result{
				Current:     "alice",
				Account:     "bob",
				SourceKind:  Local,
				SourceValue: "path/to/local/config",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := resolveLocal(tt.path, tt.gh, tt.loader)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("resolveLocal() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("resolveLocal() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("resolveLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
