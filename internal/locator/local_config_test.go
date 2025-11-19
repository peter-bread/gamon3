package locator_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/peter-bread/gamon3/internal/locator"
)

type mockLocalOS struct {
	wd      string
	wdErr   error
	home    string
	homeErr error
	files   map[string]bool
	statErr error
}

func (m mockLocalOS) Getwd() (string, error) {
	if m.wdErr != nil {
		return "", m.wdErr
	}
	return m.wd, nil
}

func (m mockLocalOS) UserHomeDir() (string, error) {
	if m.homeErr != nil {
		return "", m.homeErr
	}
	return m.home, nil
}

func (m mockLocalOS) Stat(name string) (os.FileInfo, error) {
	if m.statErr != nil {
		return nil, m.statErr
	}

	if m.files[name] {
		return nil, nil
	}

	return nil, os.ErrNotExist
}

func TestLocalConfigPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		os      locator.LocalOS
		want    string
		wantErr bool
	}{
		{
			name: "error if cannot get working directory",
			os: mockLocalOS{
				wdErr: fmt.Errorf("could not get working directory"),
			},
			wantErr: true,
		},
		{
			name: "error if cannot get user home directory",
			os: mockLocalOS{
				wd:      "/home/user/project",
				homeErr: fmt.Errorf("could not get user home directory"),
			},
			wantErr: true,
		},
		{
			name: "local config file exists in current directory",
			os: mockLocalOS{
				wd:    "/home/user/project",
				home:  "/home/user",
				files: map[string]bool{"/home/user/project/.gamon3.yaml": true},
			},
			want: "/home/user/project/.gamon3.yaml",
		},
		{
			name: "local config file exists in parent directory",
			os: mockLocalOS{
				wd:    "/home/user/project/subdir",
				home:  "/home/user",
				files: map[string]bool{"/home/user/project/.gamon3.yml": true},
			},
			want: "/home/user/project/.gamon3.yml",
		},
		{
			name: "local config file not found",
			os: mockLocalOS{
				wd:    "/home/user/project",
				home:  "/home/user",
				files: map[string]bool{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := locator.LocalConfigPath(tt.os)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LocalConfigPath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LocalConfigPath() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("LocalConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
