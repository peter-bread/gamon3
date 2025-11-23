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

package matcher

import (
	"os"
	"testing"
)

func Test_normalise(t *testing.T) {
	home, _ := os.UserHomeDir()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path string
		want string
	}{
		{
			name: "handle empty path",
			path: "",
			want: "",
		},
		{
			name: "handle path being ~",
			path: "~",
			want: home,
		},
		{
			name: "handle path being $HOME",
			path: "$HOME",
			want: home,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalise(tt.path)
			if got != tt.want {
				t.Errorf("normalise() = %v, want %v", got, tt.want)
			}
		})
	}
}
