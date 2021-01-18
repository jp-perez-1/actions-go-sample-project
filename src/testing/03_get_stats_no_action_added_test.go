package main

import (
	"testing"

	"../actions"
)

// Test if GetStats is empty if no add action were ran
func TestGetStatsNoActionAdded(t *testing.T) {

	actions := actions.CreateActionObject()
	//Should be empty as no add action was added
	if actions.GetStats() != "[]" {
		t.Errorf("Was expecting [].")
	}
}
