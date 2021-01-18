package main

import (
	"testing"

	"../actions"
)

// Test if AddAction can still proceed after trying to add erroneous JSON strings
func TestImproperAddActionJSON(t *testing.T) {

	actions := actions.CreateActionObject()

	// Test for improper formatted JSON (returns error)
	if actions.AddAction("{\"action\":\"jump, \"time\":100}") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of improper JSON input.")
	}
	// Test for empty JSON object {} (returns error)
	if actions.AddAction("{}") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of being an empty JSON input {}.")
	}
	// Test for missing the field time(returns error)
	if actions.AddAction("{\"action\":\"jump\"}") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of missing the 'time' field.")
	}
	// Test to ensure time with 0 works as Go treats 0 as the default empty value of an int (returns nil)
	if actions.AddAction("{\"action\":\"jump\", \"time\":0}") != nil {
		t.Errorf("AddAction did return an error when it shouldn't have ('time' of 0 should be valid).")
	}
	// Test to ensure empty action is treated as erroneous (returns error)
	if actions.AddAction("{\"action\":\"\", \"time\":0}") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of 'action' field being empty.")
	}
	// Test to ensure missing action field is treated as erroneous (returns error)
	if actions.AddAction("{\"time\":0}") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of input missing 'action'.")
	}
	// Test to ensure array is treated as erroneous (returns error)
	if actions.AddAction("[]") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of input being an empty array [].")
	}
	// Test to ensure blank string is treated as erroneous (returns error)
	if actions.AddAction("") == nil {
		t.Errorf("AddAction did not return an error when it should have as a result of input being empty string.")
	}
	// Test to ensure string casing doesn't matter when reading JSON
	if actions.AddAction("{\"ActIon\":\"jump\", \"time\":0}") != nil {
		t.Errorf("AddAction did return an error when it shouldn't have (casing in JSON shouldn't matter).")
	}
}
