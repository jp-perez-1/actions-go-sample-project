package main

import (
	"testing"

	actions "actions/go/src/actions"
)

// Basic example taken from the assignment prompt
func TestGetStatsBasicExample(t *testing.T) {

	actions := actions.CreateActionObject()
	var s1 string = "{\"action\":\"jump\", \"time\":100}"
	var s2 string = "{\"action\":\"run\", \"time\":75}"
	var s3 string = "{\"action\":\"jump\", \"time\":200}"
	actions.AddAction(s1)
	actions.AddAction(s2)
	actions.AddAction(s3)
	// Should result in the same array with "jump" and an avg of 150 and "run" and an avg of 75
	if actions.GetStats() != "[{\"action\":\"jump\",\"avg\":150},{\"action\":\"run\",\"avg\":75}]" {
		t.Errorf("Example failed got %s was expecting [{\"action\":\"jump\",\"avg\":150},{\"action\":\"run\",\"avg\":75}].", actions.GetStats())
	}
}
