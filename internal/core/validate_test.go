package core_test

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/core"
)

func errorStrings(errs []error) []string {
	strs := make([]string, len(errs))
	for i, e := range errs {
		strs[i] = e.Error()
	}
	return strs
}

func TestValidateMainConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cfg     *config.MainConfig
		ghHosts []string
		want    []error
	}{
		{
			name:    "empty main config and empty gh hosts should error",
			cfg:     &config.MainConfig{},
			ghHosts: []string{},
			want:    []error{fmt.Errorf("default account is required")},
		},
		{
			name: "default missing in main config but account present in gh hosts",
			cfg: &config.MainConfig{
				Accounts: map[string][]string{
					"john": {"$HOME/john"},
				},
			},
			ghHosts: []string{"john"},
			want:    []error{fmt.Errorf("default account is required")},
		},
		{
			name: "default present and authenticated, single account present in gh hosts",
			cfg: &config.MainConfig{
				Default: "john",
				Accounts: map[string][]string{
					"john": {"$HOME/john"},
				},
			},
			ghHosts: []string{"john"},
			want:    []error{},
		},
		{
			name: "default present but not in gh hosts, account present",
			cfg: &config.MainConfig{
				Default: "john",
				Accounts: map[string][]string{
					"john": {"$HOME/john"},
				},
			},
			ghHosts: []string{},
			want: []error{
				fmt.Errorf("default account %q is not authenticated with GH CLI", "john"),
				fmt.Errorf("account %q is not authenicated with GH CLI", "john"),
			},
		},
		{
			name: "multiple accounts, some missing in gh hosts",
			cfg: &config.MainConfig{
				Default: "alice",
				Accounts: map[string][]string{
					"alice": {"$HOME/alice"},
					"bob":   {"$HOME/bob"},
				},
			},
			ghHosts: []string{"alice"},
			want: []error{
				fmt.Errorf("account %q is not authenicated with GH CLI", "bob"),
			},
		},
		{
			name: "multiple accounts, all authenticated",
			cfg: &config.MainConfig{
				Default: "alice",
				Accounts: map[string][]string{
					"alice": {"$HOME/alice"},
					"bob":   {"$HOME/bob"},
				},
			},
			ghHosts: []string{"alice", "bob"},
			want:    []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.ValidateMainConfig(tt.cfg, tt.ghHosts)

			// Convert to strings as errors cannot be sorted by default.
			gotStr := errorStrings(got)
			wantStr := errorStrings(tt.want)

			// Sort error strings so they can be compared properly.
			sort.Strings(gotStr)
			sort.Strings(wantStr)

			if !slices.Equal(gotStr, wantStr) {
				t.Errorf("ValidateMainConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
