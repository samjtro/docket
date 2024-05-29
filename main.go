package main

import (
	"bufio"
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
	var Docket DOCKET
	pathBase = fmt.Sprintf("%s/.foo/docket", homeDir())

	if _, err := os.Stat(pathBase); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathBase, 0755)

		if err != nil {
			log.Fatalf(err.Error())
		}

		f, err := os.Create(pathElements + "/personal.json")

		if err != nil {
			log.Fatalf(err.Error())
		}

		f.Close()
		elementNames = append(elementNames, "personal")

		f, err = os.Create(pathElements + "/work.json")

		if err != nil {
			log.Fatalf(err.Error())
		}

		f.Close()
		elementNames = append(elementNames, "work")
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

	for _, element := range elementNames {
		// Find a better way to get the first line
		f, err := os.Open(fmt.Sprintf("%s/%s.json", pathElements, element))
		check(err)
		defer f.Close()

		fileScanner := bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)

		var line int
		for fileScanner.Scan() {
			if line == 0 {
				if fileScanner.Text() == "t" {

				}
			}
			line++
		}
	}
}

func main() {
	switch os.Args[1] {
	case "glance", "g":
		glance(os.Args[2])
	case "add", "a":
		switch os.Args[2] {
		case "event", "e":
			fifthArgument, err := strconv.Atoi(os.Args[5])
			check(err)
			sixthArgument, err := strconv.Atoi(os.Args[6])
			check(err)

			createEvent(os.Args[3], os.Args[4], fifthArgument, sixthArgument)
		case "task", "t":
			fifthArgument, err := strconv.Atoi(os.Args[5])
			check(err)

			createTask(os.Args[3], os.Args[4], fifthArgument)
		}
	case "search", "s":
		search()
	}
}
