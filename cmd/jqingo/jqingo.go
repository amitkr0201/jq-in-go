package jqingo

import (
	"bytes"
	"encoding/json"
	"os"
)

// ProcessCommand takes all arguments from the command line and does work accordingly
func ProcessCommand(compact bool, args []string) (err error) {
	input := args[0]

	output, err := formatOutput(compact, input)
	if err != nil {
		return
	}

	printOutput(output)

	return
}

func printOutput(input []byte) {
	os.Stdout.Write(input)
}

func formatOutput(compact bool, input string) (output []byte, err error) {
	var f interface{}
	output = []byte{}

	// Read as data
	if err = json.Unmarshal([]byte(input), &f); err != nil {
		return
	}

	// If compact flag is on
	if compact {
		var c []byte

		c, err = json.Marshal(f)
		if err != nil {
			return
		}
		os.Stdout.Write(c)
		buffer := new(bytes.Buffer)

		if err = json.Compact(buffer, c); err != nil {
			return
		}

		output = buffer.Bytes()
		os.Stdout.Write(output)
		return
	}

	// Get pretty-printed string
	if output, err = json.MarshalIndent(f, "", "  "); err != nil {
		return
	}

	return
}
