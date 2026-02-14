/*
Copyright © 2025 Peter Sheehan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package validate_test

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/peter-bread/gamon3/v2/internal/config"
	"github.com/peter-bread/gamon3/v2/internal/validate"
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
				fmt.Errorf("account %q is not authenticated with GH CLI", "john"),
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
				fmt.Errorf("account %q is not authenticated with GH CLI", "bob"),
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
			got := validate.ValidateMainConfig(tt.cfg, tt.ghHosts)

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

func TestValidateLocalConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cfg     *config.LocalConfig
		ghHosts []string
		wantErr bool
	}{
		{
			name:    "empty config should error",
			cfg:     &config.LocalConfig{},
			ghHosts: []string{},
			wantErr: true,
		},
		{
			name:    "account not in ghHosts should error",
			cfg:     &config.LocalConfig{Account: "john"},
			ghHosts: []string{"alice", "bob"},
			wantErr: true,
		},
		{
			name:    "account present in ghHosts should succeed",
			cfg:     &config.LocalConfig{Account: "john"},
			ghHosts: []string{"alice", "john", "bob"},
			wantErr: false,
		},
		{
			name:    "empty ghHosts with account set should error",
			cfg:     &config.LocalConfig{Account: "john"},
			ghHosts: []string{},
			wantErr: true,
		},
		{
			name:    "account with different case in ghHosts should error (case-sensitive)",
			cfg:     &config.LocalConfig{Account: "John"},
			ghHosts: []string{"john"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := validate.ValidateLocalConfig(tt.cfg, tt.ghHosts)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ValidateLocalConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ValidateLocalConfig() succeeded unexpectedly")
			}
		})
	}
}
