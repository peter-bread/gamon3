package locator

import (
	"os"

	"github.com/peter-bread/gamon3/internal/data"
)

func EnvAccount() (string, bool) {
	return os.LookupEnv(data.EnvVarAccount)
}
