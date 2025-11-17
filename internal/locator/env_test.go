package locator_test

import (
	"testing"

	"github.com/peter-bread/gamon3/internal/data"
	"github.com/peter-bread/gamon3/internal/locator"
)

func TestEnvAccount(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		env   string
		isSet bool
		want  string
		want2 bool
	}{
		{
			name:  "false if env var is unset",
			isSet: false,
			want2: false,
		},
		{
			name:  "true if env var is set but empty",
			isSet: true,
			want2: true,
		},
		{
			name:  "true if env var is set and non-empty",
			env:   "hello",
			isSet: true,
			want:  "hello",
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isSet {
				t.Setenv(data.EnvVarAccount, tt.env)
			}

			got, got2 := locator.EnvAccount()
			if got != tt.want {
				t.Errorf("EnvAccount() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("EnvAccount() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
