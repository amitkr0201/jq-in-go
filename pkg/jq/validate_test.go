package jq

import "testing"

func TestValidJSON(t *testing.T) {
	input := `{"test":"yes"}`
	valid := IsJSONValid(input)
	if !valid {
		t.Fatalf("Failed test to check JSON validity.")
	}
}
func TestInvalidJSON(t *testing.T) {
	input := `{"test":"yes",}`
	valid := IsJSONValid(input)
	if valid {
		t.Fatalf("Failed test to check JSON validity.")
	}
}
