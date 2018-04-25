package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amitkr0201/jq-in-go/pkg/cmd"
)

func main() {
	flag.Parse()
	args := flag.Args()
	err := cmd.ProcessCommand(args)

	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
}
