package config_test

import (
	"reflect"
	"testing"

	"github.com/peter-bread/gamon3/internal/config"
)

func TestLoadLocalConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.LocalConfig
		wantErr bool
	}{
		{
			name: "valid config",
			path: "testdata/local_config/valid.yaml",
			want: &config.LocalConfig{
				Account: "jane",
			},
		},
		{
			name:    "invalid config",
			path:    "testdata/local_config/invalid.yaml",
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
			got, gotErr := config.LoadLocalConfig(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadLocalConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadLocalConfig() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadLocalConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
