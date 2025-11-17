package resolve_test

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/data"
	"github.com/peter-bread/gamon3/internal/resolve"
)

type mockLocator struct {
	envAccount     string
	envFound       bool
	localPath      string
	localErr       error
	mainPath       string
	mainErr        error
	ghHostsPath    string
	ghHostsPathErr error
}

func (m mockLocator) GhHostsPath() (string, error)     { return m.ghHostsPath, m.ghHostsPathErr }
func (m mockLocator) EnvAccount() (string, bool)       { return m.envAccount, m.envFound }
func (m mockLocator) LocalConfigPath() (string, error) { return m.localPath, m.localErr }
func (m mockLocator) MainConfigPath() (string, error)  { return m.mainPath, m.mainErr }

type mockGhHosts struct {
	current string
	users   []string
}

func (m mockGhHosts) CurrentUser() string { return m.current }
func (m mockGhHosts) AllUsers() []string  { return m.users }

type mockGhHostsLoader struct {
	gh  resolve.GhHosts
	err error
}

func (m mockGhHostsLoader) Load(path string) (resolve.GhHosts, error) {
	return m.gh, m.err
}

type mockLocalLoader struct {
	cfg *config.LocalConfig
	err error
}

func (m mockLocalLoader) Load(path string) (*config.LocalConfig, error) {
	return m.cfg, m.err
}

type mockMainLoader struct {
	cfg *config.MainConfig
	err error
}

func (m mockMainLoader) Load(path string) (*config.MainConfig, error) {
	return m.cfg, m.err
}

type mockOS struct {
	cwd string
	err error
}

func (m mockOS) Getwd() (string, error) {
	return m.cwd, m.err
}

func TestResolve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		l           resolve.Locator
		ghLoader    resolve.GhHostsLoader
		localLoader resolve.LocalConfigLoader
		mainLoader  resolve.MainConfigLoader
		os          resolve.OS
		want        resolve.Result
		wantErr     bool
	}{
		// GH Hosts
		{
			name: "error if gh hosts path fails",
			l: mockLocator{
				ghHostsPathErr: fmt.Errorf("could not locate gh hosts path"),
			},
			wantErr: true,
		},
		{
			name: "error if gh loader fails",
			l:    mockLocator{},
			ghLoader: mockGhHostsLoader{
				err: fmt.Errorf("gh loader failed"),
			},
			wantErr: true,
		},

		// No methods
		{
			name: "error if no config sources found",
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainErr:  fmt.Errorf("no main config"),
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			wantErr: true,
		},

		// Env
		{
			name: "error if env var is set and empty",
			l: mockLocator{
				envAccount: "",
				envFound:   true,
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			wantErr: true,
		},
		{
			name: "error if env var is set to invalid account",
			l: mockLocator{
				envAccount: "invalid",
				envFound:   true,
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			wantErr: true,
		},
		{
			name: "env var should be used if it is the only config source",
			l: mockLocator{
				envAccount: "alice",
				envFound:   true,
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			want: resolve.Result{
				Current:     "bob",
				Account:     "alice",
				SourceKind:  resolve.Env,
				SourceValue: data.EnvVarAccount,
			},
		},
		{
			name: "env var should override local config",
			l: mockLocator{
				envAccount: "bob",
				envFound:   true,
				localPath:  "path/to/local/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			localLoader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "alice",
				},
			},
			want: resolve.Result{
				Current:     "bob",
				Account:     "bob",
				SourceKind:  resolve.Env,
				SourceValue: data.EnvVarAccount,
			},
		},
		{
			name: "env var should override main config",
			l: mockLocator{
				envAccount: "bob",
				envFound:   true,
				localErr:   fmt.Errorf("no local config"),
				mainPath:   "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "alice",
				},
			},
			want: resolve.Result{
				Current:     "bob",
				Account:     "bob",
				SourceKind:  resolve.Env,
				SourceValue: data.EnvVarAccount,
			},
		},
		{
			name: "env var should override local config and main config",
			l: mockLocator{
				envAccount: "dave",
				envFound:   true,
				localPath:  "path/to/local/config",
				mainPath:   "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob", "claire", "dave"},
				},
			},
			localLoader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "claire",
				},
			},
			mainLoader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "bob",
				},
			},
			want: resolve.Result{
				Current:     "alice",
				Account:     "dave",
				SourceKind:  resolve.Env,
				SourceValue: data.EnvVarAccount,
			},
		},

		// Local
		{
			name: "error if local config loader fails",
			l: mockLocator{
				envFound:  false,
				localPath: "path/to/local/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			localLoader: mockLocalLoader{
				err: fmt.Errorf("failed to load local config"),
			},
			wantErr: true,
		},
		{
			name: "error if local config account is empty",
			l: mockLocator{
				envFound:  false,
				localPath: "path/to/local/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			localLoader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "",
				},
			},
			wantErr: true,
		},
		{
			name: "error if local config account is invalid",
			l: mockLocator{
				envFound:  false,
				localPath: "path/to/local/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			localLoader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "invalid",
				},
			},
			wantErr: true,
		},
		{
			name: "use local config if it is the only config source",
			l: mockLocator{
				envFound:  false,
				localPath: "path/to/local/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			localLoader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "alice",
				},
			},
			want: resolve.Result{
				Current:     "bob",
				Account:     "alice",
				SourceKind:  resolve.Local,
				SourceValue: "path/to/local/config",
			},
		},
		{
			name: "local config should override main config",
			l: mockLocator{
				envFound:  false,
				localPath: "path/to/local/config",
				mainPath:  "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob", "claire"},
				},
			},
			localLoader: mockLocalLoader{
				cfg: &config.LocalConfig{
					Account: "bob",
				},
			},
			mainLoader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "claire",
				},
			},
			want: resolve.Result{
				Current:     "alice",
				Account:     "bob",
				SourceKind:  resolve.Local,
				SourceValue: "path/to/local/config",
			},
		},

		// Main
		{
			name: "error if main config loader fails",
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
				err: fmt.Errorf("main config loader failed"),
			},
			wantErr: true,
		},
		{
			name: "error main config if os.Getwd fails",
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "alice",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
				cfg: &config.MainConfig{
					Default: "alice",
				},
			},
			os: mockOS{
				cwd: "anywhere",
			},
			want: resolve.Result{
				Current:     "bob",
				Account:     "alice",
				SourceKind:  resolve.Main,
				SourceValue: "path/to/main/config",
			},
		},
		{
			name: "use main config default account if other account mappings are defined but not matched",
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			want: resolve.Result{
				Current:     "bob",
				Account:     "alice",
				SourceKind:  resolve.Main,
				SourceValue: "path/to/main/config",
			},
		},
		{
			name: "use main config account from mapping if cwd matches",
			l: mockLocator{
				envFound: false,
				localErr: fmt.Errorf("no local config"),
				mainPath: "path/to/main/config",
			},
			ghLoader: mockGhHostsLoader{
				gh: mockGhHosts{
					current: "bob",
					users:   []string{"alice", "bob"},
				},
			},
			mainLoader: mockMainLoader{
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
			want: resolve.Result{
				Current:     "bob",
				Account:     "bob",
				SourceKind:  resolve.Main,
				SourceValue: "path/to/main/config",
			},
		},
		// TODO: Possibly more, but I think I've got most cases covered.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := resolve.Resolve(tt.l, tt.ghLoader, tt.localLoader, tt.mainLoader, tt.os)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Resolve() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Resolve() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
