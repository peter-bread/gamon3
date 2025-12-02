package main

import (
	"log"
	"os"

	"github.com/peter-bread/gamon3/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := cmd.RootCmd

	os.Mkdir("./docs", 0o755)

	err := doc.GenMarkdownTree(cmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./man", 0o755)

	err = doc.GenManTree(cmd, nil, "./man")
	if err != nil {
		log.Fatal(err)
	}
}
