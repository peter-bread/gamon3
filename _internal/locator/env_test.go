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

	"github.com/peter-bread/gamon3/v2/internal/data"
	"github.com/peter-bread/gamon3/v2/internal/locator"
)

type mockEnvOS struct {
	env map[string]string
}

func (m mockEnvOS) LookupEnv(key string) (string, bool) {
	val, ok := m.env[key]
	return val, ok
}

func TestEnvAccount(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		os    locator.EnvOS
		want  string
		want2 bool
	}{
		{
			name: fmt.Sprintf("%s env var is not set", data.EnvVarAccount),
			os: mockEnvOS{
				env: map[string]string{},
			},
			want:  "",
			want2: false,
		},
		{
			name: fmt.Sprintf("%s env var is set", data.EnvVarAccount),
			os: mockEnvOS{
				env: map[string]string{data.EnvVarAccount: "hello"},
			},
			want:  "hello",
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := locator.EnvAccount(tt.os)
			if got != tt.want {
				t.Errorf("EnvAccount() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("EnvAccount() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
