package resolve

import "slices"

func IsValidGitHubAccount(account string, gh GhHosts) bool {
	return slices.Contains(gh.AllUsers(), account)
}
