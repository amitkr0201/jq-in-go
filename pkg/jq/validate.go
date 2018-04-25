package jq

import (
	"encoding/json"
	"errors"
)

// IsJSONValid tests if the provided string is a valid JSON or not.
func IsJSONValid(input string) error {
	var err error

	valid := json.Valid([]byte(input))

	if !valid {
		err := errors.New("Provided JSON is not valid")
		return err
	}
	return err
}
