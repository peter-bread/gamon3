package matcher

import (
	"os"
	"testing"
)

func Test_normalise(t *testing.T) {
	home, _ := os.UserHomeDir()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path string
		want string
	}{
		{
			name: "handle empty path",
			path: "",
			want: "",
		},
		{
			name: "handle path being ~",
			path: "~",
			want: home,
		},
		{
			name: "handle path being $HOME",
			path: "$HOME",
			want: home,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalise(tt.path)
			if got != tt.want {
				t.Errorf("normalise() = %v, want %v", got, tt.want)
			}
		})
	}
}
