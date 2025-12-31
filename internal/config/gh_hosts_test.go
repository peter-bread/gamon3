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

package config_test

import (
	"reflect"
	"slices"
	"sort"
	"testing"

	"github.com/peter-bread/gamon3/v2/internal/config"
)

func TestLoadGhHosts(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.GhHosts
		wantErr bool
	}{
		{
			name: "parse valid hosts file",
			path: "testdata/gh_hosts/valid.yaml",
			want: &config.GhHosts{
				GitHubCom: struct {
					Users map[string]any `yaml:"users"`
					User  string         `yaml:"user"`
				}{
					Users: map[string]any{
						"alice": nil,
						"bob":   nil,
					},
					User: "alice",
				},
			},
		},
		{
			name:    "error if cannot read file",
			path:    "path/to/file/that/does/not/exist",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := config.LoadGhHosts(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadGhHosts() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadGhHosts() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadGhHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGhHosts_GetCurrentUser(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		path string
		want string
	}{
		{
			name: "current user is alice",
			path: "testdata/gh_hosts/valid.yaml",
			want: "alice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := config.LoadGhHosts(tt.path)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := g.CurrentUser()
			if got != tt.want {
				t.Errorf("GetCurrentUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGhHosts_GetAllUsers(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		path string
		want []string
	}{
		{
			name: "authenticated users are bob and alice",
			path: "testdata/gh_hosts/valid.yaml",
			want: []string{"alice", "bob"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := config.LoadGhHosts(tt.path)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := g.AllUsers()
			sort.Strings(got) // Need to sort as `yaml.Unmarshal` doesn't always return in the same order.
			if !slices.Equal(got, tt.want) {
				t.Errorf("GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
