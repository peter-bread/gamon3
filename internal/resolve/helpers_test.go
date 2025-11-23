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

import "testing"

func Test_isValidGitHubAccount(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		account string
		gh      GhHosts
		want    bool
	}{
		{
			name:    "returns false when no users exist",
			account: "alice",
			gh:      mockGhHosts{users: []string{}},
			want:    false,
		},
		{
			name:    "returns false when account not in authenticated users",
			account: "alice",
			gh:      mockGhHosts{users: []string{"bob"}},
			want:    false,
		},
		{
			name:    "returns true when account is in authenticated users",
			account: "alice",
			gh:      mockGhHosts{users: []string{"alice"}},
			want:    true,
		},
		{
			name:    "returns true when account appears among multiple users",
			account: "alice",
			gh:      mockGhHosts{users: []string{"alice", "bob"}},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidGitHubAccount(tt.account, tt.gh)
			if got != tt.want {
				t.Errorf("isValidGitHubAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
