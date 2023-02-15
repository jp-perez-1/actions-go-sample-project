package main

import (
	"sync"
	"testing"

	actions "actions/go/src/actions"
)

// Test if concurrency works by running a lot of "go" commands and ensures there is no race condition / crash
func TestConcurrency(t *testing.T) {
	actions := actions.CreateActionObject()

	var s1 string = "{\"action\":\"jump\", \"time\":100}"
	var s2 string = "{\"action\":\"run\", \"time\":75}"
	var s3 string = "{\"action\":\"jump\", \"time\":200}"
	var s4 string = "{\"action\":\"duck\", \"time\":55}"
	var s5 string = "{\"action\":\"duck\", \"time\":20}"

	// Wait group to synchronize go commands
	var wg sync.WaitGroup
	i := 0

	//Run 100,000 times the same loop of the above 5 strings
	//Every 10 iterations run getStats
	//Run the functions concurrently using "go"
	for i < 100000 {
		i++
		wg.Add(1)
		go func() {
			actions.AddAction(s1)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			actions.AddAction(s2)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			actions.AddAction(s3)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			actions.AddAction(s4)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			actions.AddAction(s5)
			wg.Done()
		}()
		if i%10 == 0 {
			wg.Add(1)
			go func() {
				actions.GetStats()
				wg.Done()
			}()
		}
	}
	//wait for all processes to finish
	wg.Wait()

	// Despite all the processes running concurrently the final result should remain as below
	// If concurrency was not handled the result would likely be random or a crash might happen
	if actions.GetStats() != "[{\"action\":\"duck\",\"avg\":37.5},{\"action\":\"jump\",\"avg\":150},{\"action\":\"run\",\"avg\":75}]" {
		t.Errorf("Example failed got %s was expecting [{\"action\":\"duck\",\"avg\":37.5},{\"action\":\"jump\",\"avg\":150},{\"action\":\"run\",\"avg\":75}].", actions.GetStats())
	}
}
