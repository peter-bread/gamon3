package cmd

import (
	"flag"
	"fmt"
	"os"
	"peter-bread/gamon3/cmd/run"
	"peter-bread/gamon3/cmd/setup"
)

func usage() {
	fmt.Println("Usage: gamon3 <setup | run>")
	os.Exit(1)
}

func Execute() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	switch args[0] {

	case "setup":
		setup.SetupCmd()

	case "run":
		run.RunCmd()

	default:
		usage()
	}
}
