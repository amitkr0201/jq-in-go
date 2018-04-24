package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amitkr0201/jq-in-go/pkg/jq"
)

func main() {
	flag.Parse()
	args := flag.Args()
	result := processCommand(args)

	if result {
		fmt.Print("Valid JSON provided.")
	} else {
		fmt.Print("Provided JSON is not valid.")
	}
}

func processCommand(args []string) bool {
	if len(args) == 0 {
		fmt.Printf("No arguments provided. Nothing to do.")
		os.Exit(0)
	}
	if len(args) != 1 {
		fmt.Printf("This command can only process 1 argument. Failing...")
		os.Exit(1)
	}
	input := args[0]
	result := jq.IsJSONValid(input)
	return result
}
