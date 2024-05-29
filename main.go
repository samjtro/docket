package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	elementNames []string
	pathBase     string
	pathElements string = pathBase + "/elements"
	pathCache    string = pathBase + "/cache"
	pathConfig   string = pathBase + "/config"
)

func init() {
	// var Docket DOCKET
	pathBase = fmt.Sprintf("%s/.foo/docket", HomeDir())

	if _, err := os.Stat(pathBase); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathBase, 0755)

		if err != nil {
			log.Fatalf(err.Error())
		}

		f, err := os.Create(pathElements + "/default.json")

		if err != nil {
			log.Fatalf(err.Error())
		}

		defer f.Close()
	} else {
		filepath.WalkDir(pathElements, func(s string, d fs.DirEntry, e error) error {
			if e != nil {
				log.Fatalf(err.Error())
			}

			if filepath.Ext(d.Name()) == ".json" {
				elementNames = append(elementNames, s)
			}

			return nil
		})
	}
}

func main() {
	switch os.Args[1] {
	case "glance":
		Glance(os.Args[2])
	case "add":
		switch os.Args[2] {
		case "e", "event":
			fifthArgument, err := strconv.Atoi(os.Args[5])
			Check(err)
			sixthArgument, err := strconv.Atoi(os.Args[6])
			Check(err)

			CreateEvent(os.Args[3], os.Args[4], fifthArgument, sixthArgument)
		case "t", "task":
			fifthArgument, err := strconv.Atoi(os.Args[5])
			Check(err)

			CreateTask(os.Args[3], os.Args[4], fifthArgument)
		}
	case "search":
		Search()
	}
}
