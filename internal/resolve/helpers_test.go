package resolve

import "testing"

func Test_isValidGitHubAccount(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		account string
		gh      GhHosts
		want    bool
	}{
		{
			name:    "returns false when no users exist",
			account: "alice",
			gh:      mockGhHosts{users: []string{}},
			want:    false,
		},
		{
			name:    "returns false when account not in authenticated users",
			account: "alice",
			gh:      mockGhHosts{users: []string{"bob"}},
			want:    false,
		},
		{
			name:    "returns true when account is in authenticated users",
			account: "alice",
			gh:      mockGhHosts{users: []string{"alice"}},
			want:    true,
		},
		{
			name:    "returns true when account appears among multiple users",
			account: "alice",
			gh:      mockGhHosts{users: []string{"alice", "bob"}},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidGitHubAccount(tt.account, tt.gh)
			if got != tt.want {
				t.Errorf("isValidGitHubAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
