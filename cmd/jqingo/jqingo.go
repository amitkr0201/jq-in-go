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
	output = []byte{}

	// If compact flag is on
	if compact {
		output, err = compactOutput(input)
		return
	}

	var f interface{}

	// Read as data
	if err = json.Unmarshal([]byte(input), &f); err != nil {
		return
	}

	// MarshalIndent can't throw error because
	// f is a proper JSON from Unmarshal earlier
	output, _ = json.MarshalIndent(f, "", "  ")

	return
}

func compactOutput(input string) (output []byte, err error) {
	var f interface{}
	var c []byte

	// Read as data
	if err = json.Unmarshal([]byte(input), &f); err != nil {
		return
	}

	// Marshal can't throw error because
	// f is a proper JSON from Unmarshal earlier
	c, _ = json.Marshal(f)

	buffer := new(bytes.Buffer)

	// c is proper JSON from Marshal above
	_ = json.Compact(buffer, c)

	output = buffer.Bytes()
	return
}
