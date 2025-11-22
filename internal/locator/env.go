package locator

import (
	"github.com/peter-bread/gamon3/internal/data"
)

type EnvOS interface {
	LookupEnv(key string) (string, bool)
}

// EnvAccount checks if the $GAMON3_ACCOUNT environment variable is set.
// If set, return its value and true, otherwise return an empty string and
// false.
func EnvAccount(os EnvOS) (string, bool) {
	return os.LookupEnv(data.EnvVarAccount)
}
