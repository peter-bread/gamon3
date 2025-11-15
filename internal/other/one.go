package other

import (
	"slices"

	"github.com/peter-bread/gamon3/internal/config"
)

func IsValidGitHubAccount(account string, gh *config.GhHosts) bool {
	return slices.Contains(gh.AllUsers(), account)
}
