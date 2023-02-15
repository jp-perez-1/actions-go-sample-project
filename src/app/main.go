package main

import (
	"fmt"
	"sync"

	actions "actions/go/src/actions"
)

// This is just a sample use of the package
// Overwrite or copy and use it as a sample, see the README.md for more details regarding usage
func main() {

	actions := actions.CreateActionObject()

	var s1 string = "{\"action\":\"jump\", \"time\":100}"
	var s2 string = "{\"action\":\"run\", \"time\":75}"
	var s3 string = "{\"action\":\"jump\", \"time\":200}"
	var s4 string = "{\"action\":\"duck\", \"time\":55}"
	var s5 string = "{\"action\":\"duck\", \"time\":20}"

	// Wait group to synchronize go commands
	var wg sync.WaitGroup
	i := 0

	//Run 100000 times the same loop of the above 5 strings
	//Every 10 run getStats
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
	fmt.Printf(actions.GetStats())
}
