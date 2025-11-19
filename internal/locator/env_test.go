package locator_test

import (
	"fmt"
	"testing"

	"github.com/peter-bread/gamon3/internal/data"
	"github.com/peter-bread/gamon3/internal/locator"
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
