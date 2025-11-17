package resolve

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/config"
)

func Test_resolveMain(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		gh      GhHosts
		loader  MainConfigLoader
		os      OS
		want    Result
		wantErr bool
	}{
		{
			name: "error if main config loader fails",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				err: fmt.Errorf("main config loader failed"),
			},
			wantErr: true,
		},
		{
			name: "error main config if os.Getwd fails",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "bob",
				},
			},
			os: mockOS{
				err: fmt.Errorf("os.Getwd failed"),
			},
			wantErr: true,
		},
		{
			name: "error if main config matcher fails due to empty default account",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "",
				},
			},
			os: mockOS{
				cwd: "some/path",
			},
			wantErr: true,
		},
		{
			name: "error if main config matcher fails due to empty account key",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "bob",
					Accounts: map[string][]string{
						"": {"/some/path"},
					},
				},
			},
			os: mockOS{
				cwd: "/some/path",
			},
			wantErr: true,
		},
		{
			name: "error if main config default account is selected and invalid",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "invalid",
				},
			},
			os: mockOS{
				cwd: "anywhere",
			},
			wantErr: true,
		},
		{
			name: "error if main config account from mapping is selected and invalid",
			gh: mockGhHosts{
				current: "alice",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "bob",
					Accounts: map[string][]string{
						"invalid": {"/some/path"},
					},
				},
			},
			os: mockOS{
				cwd: "/some/path",
			},
			wantErr: true,
		},
		{
			name: "use main config default account if no other account mappings are defined",
			path: "path/to/main/config",
			gh: mockGhHosts{
				current: "bob",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "alice",
				},
			},
			os: mockOS{
				cwd: "anywhere",
			},
			want: Result{
				Current:     "bob",
				Account:     "alice",
				SourceKind:  Main,
				SourceValue: "path/to/main/config",
			},
		},
		{
			name: "use main config default account if other account mappings are defined but not matched",
			path: "path/to/main/config",
			gh: mockGhHosts{
				current: "bob",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "alice",
					Accounts: map[string][]string{
						"bob": {"/use/bob/here"},
					},
				},
			},
			os: mockOS{
				cwd: "/not/there",
			},
			want: Result{
				Current:     "bob",
				Account:     "alice",
				SourceKind:  Main,
				SourceValue: "path/to/main/config",
			},
		},
		{
			name: "use main config account from mapping if cwd matches",
			path: "path/to/main/config",
			gh: mockGhHosts{
				current: "bob",
				users:   []string{"alice", "bob"},
			},
			loader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "alice",
					Accounts: map[string][]string{
						"bob": {"/use/bob/here"},
					},
				},
			},
			os: mockOS{
				cwd: "/use/bob/here/and/here",
			},
			want: Result{
				Current:     "bob",
				Account:     "bob",
				SourceKind:  Main,
				SourceValue: "path/to/main/config",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := resolveMain(tt.path, tt.gh, tt.loader, tt.os)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("resolveMain() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("resolveMain() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("resolveMain() = %v, want %v", got, tt.want)
			}
		})
	}
}
