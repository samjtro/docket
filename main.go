package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	Docket DOCKET
	// pathCache    string = pathBase + "/cache"
	// pathConfig   string = pathBase + "/config"
)

func init() {
	pathBase := fmt.Sprintf("%s/.docket", homeDir())
	pathDocket := pathBase + "/docket.json"

	// if docket hasn't been run before, create create /home/{user}/.docket/docket.json
	// else read in docket from /home/{user}/.docket/docket.json
	if _, err := os.Stat(pathBase); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathBase, os.ModePerm)
		check(err)
		f, err := os.Create(pathDocket)
		check(err)
		f.Close()
	} else {
		docketBytes, err := os.ReadFile(pathDocket)
		check(err)
		err = json.Unmarshal(docketBytes, &Docket)
		check(err)
	}
}

func main() {
	switch os.Args[1] {
	case "glance", "g":
		Docket.glance(os.Args[2])
	case "create", "c":
		switch os.Args[2] {
		case "element", "e":
			Docket.createElement(os.Args[3])
		case "task", "t":
			fifthArgument, err := strconv.Atoi(os.Args[5])
			check(err)
			Docket.createTask(os.Args[3], os.Args[4], fifthArgument)
		case "goal", "g":
			fifthArgument, err := strconv.Atoi(os.Args[5])
			check(err)
			Docket.createGoal(os.Args[3], os.Args[4], fifthArgument)
		case "milestone", "m":
			Docket.createMilestone(os.Args[3], os.Args[4])
		}
	case "search", "s":
		Docket.search()
	}
	Docket.flush()
}
