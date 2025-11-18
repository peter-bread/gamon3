package config_test

import (
	"reflect"
	"testing"

	"github.com/peter-bread/gamon3/internal/config"
)

func TestLoadMainConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.MainConfig
		wantErr bool
	}{
		{
			name: "valid config",
			path: "testdata/main_config/valid.yaml",
			want: &config.MainConfig{
				Default: "john",
				Accounts: map[string][]string{
					"steve": {"$HOME/steve"},
				},
			},
		},
		{
			name:    "invalid config",
			path:    "testdata/main_config/invalid.yaml",
			wantErr: true,
		},
		{
			name:    "error if cannot read file",
			path:    "path/to/file/that/does/not/exist",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := config.LoadMainConfig(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadMainConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadMainConfig() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadMainConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
