package cmd

import (
	"peter-bread/gamon3/cmd/run"
	"peter-bread/gamon3/cmd/setup"
)

func Execute() {
	setup.Setup()
	run.Run()
}
