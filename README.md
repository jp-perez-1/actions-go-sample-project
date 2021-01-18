# JumpCloud Interview Assignment (Backend)

## Installation

Make sure to have Go (Golang) downloaded and installed follow instructions at https://golang.org/dl/
As of the writing of this README go1.15.6 is being used in particular.

## Structure
File structure is based on go's suggested structure but only /src has anything for now

In /src there are three folders:

/actions is the package containing the implementation of the solution for the given project
/app is a package with main.go which is a sample use case of the /actions package
/testing is a testing package filled with unit tests to test certain scenarios of use of the /actions package

## Usage
This project has a file under /src/actions/actions.go to use it do the following:

1. Import the file to a go file using absolute or relative pathing see /src/app/main.go for example of importing at the top of the file
2. Create an instance of a Actions object (actions := actions.CreateActionObject())
3. Now on that instance of the Actions object you can run AddAction(newActionJSON string) and GetStats() detailed below
4. AddAction(newActionJSON string): this takes a JSON string in the form of {"action":"jump", "time":100} and processes it to be used in GetStats
5. GetStats(): this returns JSON string in the form of an array [{"action":"jump", "avg":150},{"action":"run", "avg":75}] which takes all prior AddAction commands and averages out the time per action

In /src/app/main.go you can see an example usage of the /src/actions package
To run it go to a terminal and change directory to /src/app feel free to run the code with "go [filename].go" in a terminal

To run tests go to the terminal and change directory to [Path to Project]/src/testing
There run "go test -v" and you'll see all the premade test run successfully. Inspect the code in /src/testing to see all the test scenario attempted
In that same folder you can add you own test go file make sure it ends with _test.go and doesn't duplicate any other test function name

## Future Considerations / Limitations
- Could consider using a database and persistent memory instead of in memory use would have to rework AddAction to store to a database and GetStats to obtain data from the database. One advantage of this is the database itself would handle conccurency issues given the ACID properties.
- Instead of a library class the actions.go file could be integrated / implemented to a server such that an outside client sends calls to the server to AddAction and GetStats, would likely use a database as well, as was detailed in the prior point.
- GetStats is returned ascendingly with respect to the action name