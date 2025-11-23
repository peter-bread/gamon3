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
	"os"
	"testing"

	"github.com/peter-bread/gamon3/internal/locator"
)

type mockMainOS struct {
	files     map[string]bool
	env       map[string]string
	cfgDirErr error
}

func (m mockMainOS) Stat(name string) (os.FileInfo, error) {
	if m.files[name] {
		return nil, nil
	}
	return nil, os.ErrNotExist
}

func (m mockMainOS) UserConfigDir() (string, error) { return "/mock/home/.config", m.cfgDirErr }

func (m mockMainOS) LookupEnv(key string) (string, bool) {
	val, ok := m.env[key]
	return val, ok
}

func TestMainConfigPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		os      locator.MainOS
		want    string
		wantErr bool
	}{
		{
			name: "error if main config file does not exist",
			os: mockMainOS{
				files: map[string]bool{},
				env:   map[string]string{},
			},
			wantErr: true,
		},
		{
			name: "error if cannot find user config dir",
			os: mockMainOS{
				files:     map[string]bool{},
				env:       map[string]string{},
				cfgDirErr: fmt.Errorf("could not find user config dir"),
			},
			wantErr: true,
		},
		{
			name: "config file exists in user config dir",
			os: mockMainOS{
				files: map[string]bool{"/mock/home/.config/gamon3/config.yaml": true},
				env:   map[string]string{},
			},
			want: "/mock/home/.config/gamon3/config.yaml",
		},
		{
			name: "find config file in GAMON3_CONFIG_DIR if it is set",
			os: mockMainOS{
				files: map[string]bool{"/mock/home/.config/gamon3/config.yml": true},
				env:   map[string]string{"GAMON3_CONFIG_DIR": "/mock/home/.config/gamon3"},
			},
			want: "/mock/home/.config/gamon3/config.yml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := locator.MainConfigPath(tt.os)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("MainConfigPath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("MainConfigPath() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("MainConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
