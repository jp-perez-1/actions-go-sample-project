package main

import (
	"testing"

	"../actions"
)

// Test if AddAction can still proceed after trying to add an erroneous JSON string
func TestImproperAddActionJSON(t *testing.T) {

	actions := actions.CreateActionObject()
	//Note it is missing the end quote in jump
	var s1 string = "{\"action\":\"jump, \"time\":100}"

	//Note the function returns an error
	actions.AddAction(s1)
	//Should be empty as the JSON string added was not a proper format
	if actions.GetStats() != "[]" {
		t.Errorf("Add action worked when it should not have added the erroneous JSON.")
	}
}
