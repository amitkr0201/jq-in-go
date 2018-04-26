package jqingo

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
)

// ProcessCommand takes all arguments from the command line and does work accordingly
func ProcessCommand(compact bool, args []string) (err error) {
	input := args[0]
	if result := json.Valid([]byte(input)); !result {
		return errors.New("Provided JSON is not valid")
	}

	output, err := formatOutput(compact, input)
	if err != nil {
		return
	}

	if err = printOutput(output); err != nil {
		return
	}
	return
}

func printOutput(input []byte) (err error) {
	_, err = os.Stdout.Write(input)
	return
}

func formatOutput(compact bool, input string) (b []byte, err error) {
	var f interface{}
	b = []byte{}

	// Read as data
	if err = json.Unmarshal([]byte(input), &f); err != nil {
		return
	}

	if compact {
		var c []byte

		// Get compact string
		c, err = json.Marshal(f)
		if err != nil {
			return
		}
		os.Stdout.Write(c)
		buffer := new(bytes.Buffer)

		if err = json.Compact(buffer, c); err != nil {
			return
		}

		b = buffer.Bytes()
		os.Stdout.Write(b)
		return
	}

	// Get pretty-printed string
	if b, err = json.MarshalIndent(f, "", "  "); err != nil {
		return
	}

	return
}
