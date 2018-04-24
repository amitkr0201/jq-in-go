package main

import "testing"

func TestOutputIsRight(t *testing.T) {
	expected := "I am printed."
	getting := getTextToPrint()
	if expected != getting {
		t.Fatalf("Not getting expected result. Expected: %s, Got: %s", expected, getting)
	}

	notExpected := "Unexpected."
	getting = getTextToPrint()
	if notExpected == getting {
		t.Fatalf("Getting unexpected result. Not expected %s but got it.", notExpected)
	}
}
