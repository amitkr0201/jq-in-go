package jqingo

import (
	"encoding/json"
	"fmt"
	"strings"
)

func getSelectedJSON(selector, input string) (selectedJSON string, err error) {
	var f interface{}
	splitSelctor := strings.Split(selector, ".")

	if err = json.Unmarshal([]byte(input), &f); err != nil {
		return
	}

	for i, s := range splitSelctor {
		if i != 0 {
			fmt.Printf("Selector: %s", s)
			f = applySelector(f, s)
		}
	}
	selectedJSON = interface2String(f)

	return
}

func applySelector(input interface{}, selector string) (output interface{}) {
	// if selector is blank string it means '.' is the selector
	if selector != "" {
		fmt.Printf("apply %s", selector)
		m := input.(map[string]interface{})
		output = m[selector]
		return
	}
	output = input
	return
}

func interface2String(f interface{}) (output string) {
	o, _ := json.Marshal(f)
	output = string(o)
	return
}
