# Actions Go Sample Project

## Description

A project to showcase the use of Go to make a simple class "actions.go" that is capable of adding actions in the form of JSON {"action":"jump", "time":100}, {"action":"run", "time":75}, {"action":"jump", "time":200}. 

And then returning the average of each actions time in JSON [{"action":"jump", "avg":150},{"action":"run", "avg":75}].

## Installation

Make sure to have Go (Golang) downloaded and installed follow instructions at https://golang.org/dl/

As of the writing of this README Go v1.20.1 is being used. Older or newer version might work however.

(May need to run `go mod init actions/go` in terminal in project root, but not sure...)

## Structure
File structure is based on Go's suggested structure but only /src has anything for now

In /src there are three folders:

- /actions is the package containing the implementation of the solution for the given assignment
- /app is a package with main.go which is a sample use case of the /actions package
- /testing is a testing package filled with unit tests to test certain scenarios of the main /actions package

## Usage (General)
This project has a file under /src/actions/actions.go to use it do the following:

1. Import the file to a Go file using absolute or relative pathing see /src/app/main.go for example of importing at the top of the file
2. Create an instance of an Actions object (actions := actions.CreateActionObject())
3. Now on that instance of the Actions object you can run AddAction(newActionJSON string) and GetStats() detailed below
4. AddAction(newActionJSON string): this takes a JSON string in the form of {"action":"jump", "time":100} and processes it to be used in GetStats
5. GetStats(): this returns JSON string in the form of an array [{"action":"jump", "avg":150},{"action":"run", "avg":75}] which takes all prior AddAction commands and averages out the time per action

In /src/app/main.go you can see an example usage of the /src/actions package

To run it go to a terminal and change directory to /src/app feel free to run the code with "go [filename].go" ("go main.go")

## Usage (Testing)

- To run tests go to the terminal and change directory to /src/testing
- There run "go test -v" and you'll see all the premade test run successfully. Inspect the code in /src/testing to see all the test scenario attempted
- In that same folder you can add you own test Go file, make sure it ends with _test.go and doesn't duplicate any other test function name

## Future Considerations / Limitations / Misc. Notes

- Could consider using a database and persistent memory instead of in memory use. Would have to rework AddAction to store to a database and GetStats to obtain data from the database. One advantage of this is the database itself would handle conccurency issues given the ACID properties.
- Instead of a library class the actions.go file could be integrated / implemented to a server such that an outside client sends calls to the server to AddAction and GetStats, would likely use a database as well, as was detailed in the prior point.
- GetStats is returned ascendingly with respect to the action name
- The project uses mutexes to allow for concurrency see /src/testing/05_concurrency_test.go for an example of it being used concurrently
