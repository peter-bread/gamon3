package config_test

import (
	"reflect"
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
			name: "load good gh hosts file",
			path: "./testdata/gh_hosts/good.yml",
			want: &config.GhHosts{
				"github.com": config.GhHost{
					Users: map[string]any{
						"john": nil,
					},
					User: "john",
				},
			},
		},
		{
			name:    "error if gh hosts is missing user field",
			path:    "./testdata/gh_hosts/missing_user.yml",
			wantErr: true,
		},
		{
			name:    "error if gh hosts is missing users field",
			path:    "./testdata/gh_hosts/missing_users.yml",
			wantErr: true,
		},
		{
			name:    "error if gh hosts is missing github.com field",
			path:    "./testdata/gh_hosts/missing_github.yml",
			wantErr: true,
		},
		{
			name:    "error if gh hosts file does not exist",
			path:    "./path/to/file/that/does/not/exist",
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

func TestLoadLocalConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.LocalConfig
		wantErr bool
	}{
		{
			name: "load good local config file",
			path: "./testdata/local_config/good.yml",
			want: &config.LocalConfig{
				Account: "john",
			},
		},
		{
			name:    "error if local config file contains invalid yaml syntax",
			path:    "./testdata/local_config/invalid.yml",
			wantErr: true,
		},
		{
			name:    "error if local config file is missing required account field",
			path:    "./testdata/local_config/missing.yml",
			wantErr: true,
		},
		{
			name:    "error if local config file contains additional unknown field",
			path:    "./testdata/local_config/unknown.yml",
			wantErr: true,
		},
		{
			name:    "error if local config file does not exist",
			path:    "./path/to/file/that/does/not/exist",
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

func TestLoadMainConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.MainConfig
		wantErr bool
	}{
		{
			name: "load good main config file with only default field",
			path: "./testdata/main_config/default_only.yml",
			want: &config.MainConfig{
				Default: "john",
			},
		},
		{
			name: "load good main config file with default and accounts fields",
			path: "./testdata/main_config/default_and_accounts.yml",
			want: &config.MainConfig{
				Default: "john",
				Accounts: map[string][]string{
					"jane": {"foo"},
				},
			},
		},
		{
			name:    "error if missing required default field",
			path:    "./testdata/main_config/missing_default.yml",
			wantErr: true,
		},
		{
			name:    "error if main config file contains invalid yaml syntax",
			path:    "./testdata/main_config/invalid.yml",
			wantErr: true,
		},
		{
			name:    "error if main config file does not exist",
			path:    "./path/to/file/that/does/not/exist",
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
