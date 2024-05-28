package main

import (
	"os/exec"
)

var (
	elementNames string
	pathBase string = "~/.docket"
	pathElements string = pathBase + "/elements"
	pathCache string = pathBase + "/cache"
	pathConfig string = pathBase + "/config"
)

func init() {
	var Docket DOCKET

	if _, err := os.Stat(pathBase); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathBase, 0755)

		if err != nil {
			log.Fatalf(err.Error())
		}

		f, err := os.Create(pathElements + "/main.json")

		if err != nil {
			log.Fatalf(err.Error())
		}

		defer f.Close()
	} else {
		filepath.WalkDir(pathElements, func(s string, d fs.DirEntry, e error) error {
			if e != nil {
				log.Fatalf(err.Error())
			}

			if filepath.Ext(d.Name()) == ext {
				elementNames = append(elementNames, s)
			}

			return nil
		})
	}
}

func main() {
	if os.Args[1] == "glance" {
		Glance(os.Args[2])
	} else if os.Args[1] == "add" {
		if os.Args[2] == "e" || "event" {
			CreateEvent()
		} else if os.Args[2] == "t" || "task" {
			CreateTask()
		}
	} else if os.Args[1] == "search" {
		Search()
	}
}
