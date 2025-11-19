package locator

import (
	"github.com/peter-bread/gamon3/internal/data"
)

type EnvOS interface {
	LookupEnv(key string) (string, bool)
}

func EnvAccount(os EnvOS) (string, bool) {
	return os.LookupEnv(data.EnvVarAccount)
}
