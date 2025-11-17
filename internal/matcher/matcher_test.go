package matcher_test

import (
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
		dir            string
		accounts       map[string][]string
		defaultAccount string
		want           string
		wantErr        bool
	}{
		{
			name:           "returns default when no accounts match",
			dir:            "/projects/unknown",
			accounts:       map[string][]string{"steve": {"/home/steve/work"}},
			defaultAccount: "john",
			want:           "john",
		},
		{
			name:           "matches correct account based on path prefix",
			dir:            "/home/alice/dev/app",
			accounts:       map[string][]string{"alice": {"/home/alice"}, "bob": {"/home/bob"}},
			defaultAccount: "bob",
			want:           "alice",
		},
		{
			name:           "matches using ~ expansion",
			dir:            filepath.Join(home, "projects/tool"),
			accounts:       map[string][]string{"alice": {"~/projects"}},
			defaultAccount: "bob",
			want:           "alice",
		},
		{
			name:           "matches using env var expansion",
			dir:            filepath.Join(home, "work"),
			accounts:       map[string][]string{"alice": {"$HOME/work"}},
			defaultAccount: "bob",
			want:           "alice",
		},
		{
			name:           "returns default when accounts map is empty",
			dir:            "/anything",
			accounts:       map[string][]string{},
			defaultAccount: "steve",
			want:           "steve",
		},
		{
			name:           "error when default account is empty and no match",
			dir:            "/no/match",
			accounts:       map[string][]string{"alice": {"/some/path"}},
			defaultAccount: "",
			wantErr:        true,
		},
		{
			name:           "error when account key is empty",
			dir:            "/foo",
			accounts:       map[string][]string{"": {"/foo"}},
			defaultAccount: "fallback",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := matcher.MatchAccount(tt.dir, tt.accounts, tt.defaultAccount)
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
