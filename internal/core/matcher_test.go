package core_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/peter-bread/gamon3/internal/core"
)

func TestMatchAccount(t *testing.T) {
	home, _ := os.UserHomeDir()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dir            string
		accounts       map[string][]string
		defaultAccount string
		want           string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.MatchAccount(tt.dir, tt.accounts, tt.defaultAccount)
			if got != tt.want {
				t.Errorf("MatchAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
