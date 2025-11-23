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
