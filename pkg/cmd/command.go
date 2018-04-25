package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/amitkr0201/jq-in-go/pkg/jq"
)

// ProcessCommand takes all arguments from the command line and does work accordingly
func ProcessCommand(args []string) error {
	var err error

	if len(args) == 0 {
		fmt.Printf("No arguments provided. Nothing to do.")
		os.Exit(0)
	}
	if len(args) != 1 {
		// fmt.Printf("This command can only process 1 argument. Failing...")
		return errors.New("This command can only process 1 argument")
	}

	input := args[0]
	err = jq.IsJSONValid(input)

	if err != nil {
		return err
	}

	err = printOutput(input)

	if err != nil {
		return err
	}
	return err
}

func printOutput(input string) error {
	var err error
	var f interface{}

	// Read as data
	err = json.Unmarshal([]byte(input), &f)
	if err != nil {
		return err
	}

	// Get pretty-printed string
	b, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	os.Stdout.Write(b)
	return err
}
