package jq

import (
	"encoding/json"
)

// IsJSONValid tests if the provided string is a valid JSON or not.
func IsJSONValid(input string) bool {
	return json.Valid([]byte(input))
}
