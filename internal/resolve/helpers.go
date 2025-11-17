package resolve

import "slices"

func isValidGitHubAccount(account string, gh GhHosts) bool {
	return slices.Contains(gh.AllUsers(), account)
}
