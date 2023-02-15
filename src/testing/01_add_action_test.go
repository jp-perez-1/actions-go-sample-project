package main

import (
	"testing"

	actions "actions/go/src/actions"
)

// Tests if the add action works, verified by checking GetStats
func TestAddAction(t *testing.T) {

	actions := actions.CreateActionObject()

	var s1 string = "{\"action\":\"jump\", \"time\":100}"
	actions.AddAction(s1)
	if actions.GetStats() != "[{\"action\":\"jump\",\"avg\":100}]" {
		t.Errorf("Add action failed to add the action jump with its time of 100.")
	}
}
