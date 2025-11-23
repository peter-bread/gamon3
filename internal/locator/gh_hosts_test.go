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
