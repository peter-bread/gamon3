package matcher_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/peter-bread/gamon3/internal/matcher"
)

type mockFP struct {
	AbsPath  string
	AbsError error
}

func (m mockFP) Abs(path string) (string, error) {
	if m.AbsError != nil {
		return "", m.AbsError
	}
	return m.AbsPath, nil
}

func TestMatchAccount(t *testing.T) {
	home, _ := os.UserHomeDir()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fp             matcher.FP
		dir            string
		accounts       map[string][]string
		defaultAccount string
		want           string
		wantErr        bool
	}{
		{
			name:           "returns default when no accounts match",
			fp:             mockFP{AbsPath: "/projects/unkown"},
			dir:            "/projects/unknown",
			accounts:       map[string][]string{"steve": {"/home/steve/work"}},
			defaultAccount: "john",
			want:           "john",
		},
		{
			name:           "matches correct account based on path prefix",
			fp:             mockFP{AbsPath: "/home/alice/dev/app"},
			dir:            "/home/alice/dev/app",
			accounts:       map[string][]string{"alice": {"/home/alice"}, "bob": {"/home/bob"}},
			defaultAccount: "bob",
			want:           "alice",
		},
		{
			name:           "matches using ~ expansion",
			fp:             mockFP{AbsPath: filepath.Join(home, "projects/tool")},
			dir:            filepath.Join(home, "projects/tool"),
			accounts:       map[string][]string{"alice": {"~/projects"}},
			defaultAccount: "bob",
			want:           "alice",
		},
		{
			name:           "matches using env var expansion",
			fp:             mockFP{AbsPath: filepath.Join(home, "work")},
			dir:            filepath.Join(home, "work"),
			accounts:       map[string][]string{"alice": {"$HOME/work"}},
			defaultAccount: "bob",
			want:           "alice",
		},
		{
			name:           "returns default when accounts map is empty",
			fp:             mockFP{AbsPath: "/anything"},
			dir:            "/anything",
			accounts:       map[string][]string{},
			defaultAccount: "steve",
			want:           "steve",
		},
		{
			name:           "error when default account is empty and no match",
			fp:             mockFP{AbsPath: "/no/match"},
			dir:            "/no/match",
			accounts:       map[string][]string{"alice": {"/some/path"}},
			defaultAccount: "",
			wantErr:        true,
		},
		{
			name:           "error when account key is empty",
			fp:             mockFP{AbsPath: "/foo"},
			dir:            "/foo",
			accounts:       map[string][]string{"": {"/foo"}},
			defaultAccount: "fallback",
			wantErr:        true,
		},
		{
			name:           "error if absolute filepath fails",
			fp:             mockFP{AbsError: fmt.Errorf("failed to determine absolute path")},
			dir:            "/foo",
			accounts:       map[string][]string{},
			defaultAccount: "alice",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := matcher.MatchAccount(tt.fp, tt.dir, tt.accounts, tt.defaultAccount)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("MatchAccount() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("MatchAccount() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("MatchAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
