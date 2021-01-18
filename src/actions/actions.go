package actions

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"
	"sync"
)

/*
	Actions primary object Actions has AddAction and GetStats as methods added to it see below
	Has a mutex embedded to it as to sync reads/writes into the actionData
*/
type Actions struct {
	sync.Mutex
	actionData map[string]ActionData
}

/*
	ActionInstance an instance of an Action with its name "Action" and how long it took "Time"
	Used to extract data from JSON string in AddAction
*/
type ActionInstance struct {
	Action string
	Time   int
}

/*
	ActionData aggregated Action Data Struct used with a map, "Total" is how much time overall the action took over "Count" runs
	Used to help calculate Stats
*/
type ActionData struct {
	Total int
	Count int
}

/*
	ActionStats Stat of an Action used to facilitate the production of the stats as an output JSON
	json:"[name]"` ensure that when using JSON marshal it will rename it from Action -> action, Avg -> avg
*/
type ActionStats struct {
	Action string  `json:"action"`
	Avg    float64 `json:"avg"`
}

/*
	AddAction part of the Actions struct
	This adds data from the input newActionJSON to the actionData member of Actions
*/
func (a *Actions) AddAction(newActionJSON string) error {

	var newActionInstance ActionInstance
	// Decode JSON string if it does not decode properly return an error detailing as such.
	// Check to ensure JSON string has the 'action' and 'time' fields
	if json.Unmarshal([]byte(newActionJSON), &newActionInstance) != nil {
		return errors.New("improper JSON format")
	} else if newActionInstance.Action == "" {
		return errors.New("imporper JSON format missing action field")
	} else if newActionInstance.Time == 0 && !strings.Contains(newActionJSON, "\"time\"") {
		return errors.New("imporper JSON format missing time field")
	}

	// Lock as data is about to be read and written to
	a.Lock()

	// If the action was previously recorded add to existing data else add it to the record for the first time
	if currentAction, ok := a.actionData[newActionInstance.Action]; ok {
		currentAction.Total = currentAction.Total + newActionInstance.Time
		currentAction.Count++
		a.actionData[newActionInstance.Action] = currentAction
	} else {
		a.actionData[newActionInstance.Action] = ActionData{Total: newActionInstance.Time, Count: 1}
	}

	// Unlock as data write is finished
	a.Unlock()
	return nil
}

/*
	GetStats part of the Actions struct
	Function to retrieve stats from previously added actions (average time each action takes)
*/
func (a *Actions) GetStats() string {
	actionStat := []ActionStats{}

	// Lock as data is about to be read
	a.Lock()

	// Dump all the keys in actionData to be able to sort them
	keys := make([]string, 0, len(a.actionData))
	for k := range a.actionData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Append to actionStat array the processed actionData ordered by the action's name (the key) ascendingly
	for _, k := range keys {
		v := a.actionData[k]
		actionStat = append(actionStat, ActionStats{k, float64(v.Total) / float64(v.Count)})
	}

	// Encode the data and return the JSON encoded data
	bytes, err := json.Marshal(actionStat)
	if err != nil {
		a.Unlock()
		return "[]"
	}

	a.Unlock()
	return string(bytes)

}

/*
	CreateActionObject construct a new instance of the action object
*/
func CreateActionObject() Actions {
	actions := Actions{}
	actions.actionData = map[string]ActionData{}
	return actions
}
