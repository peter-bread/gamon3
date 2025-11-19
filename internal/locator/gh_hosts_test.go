package locator_test

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/locator"
)

type mockGhOS struct {
	env       map[string]string
	cfgDirErr error
}

func (m mockGhOS) LookupEnv(key string) (string, bool) {
	val, ok := m.env[key]
	return val, ok
}
func (m mockGhOS) UserConfigDir() (string, error) { return "/mock/home/.config", m.cfgDirErr }

func TestGhHostsPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		os      locator.GhOS
		want    string
		wantErr bool
	}{
		{
			name: "return path when GH_CONFIG_DIR is set",
			os: mockGhOS{
				env: map[string]string{"GH_CONFIG_DIR": "/mock/home/.config/gh"},
			},
			want: "/mock/home/.config/gh/hosts.yml",
		},
		{
			name: "error if cannot find user config dir",
			os: mockGhOS{
				cfgDirErr: fmt.Errorf("could not find user config dir"),
			},
			wantErr: true,
		},
		{
			name: "return path when user config dir is found",
			os:   mockGhOS{},
			want: "/mock/home/.config/gh/hosts.yml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := locator.GhHostsPath(tt.os)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GhHostsPath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GhHostsPath() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GhHostsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
