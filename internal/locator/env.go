package locator

import "os"

func EnvAccount() (string, bool) {
	return os.LookupEnv("GAMON3_ACCOUNT")
}
