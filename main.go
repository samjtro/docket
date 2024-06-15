package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"
)

var (
	Docket DOCKET
	// pathCache    string = pathBase + "/cache"
	// pathConfig   string = pathBase + "/config"
)

type DOCKET struct {
	Elements []ELEMENT
}

type ELEMENT struct {
	ID         string
	Name       string // One word descriptive title, e.g. finishTaskOne or finish_task_one
	Tasks      []TASK
	Goals      []GOAL
	Milestones []MILESTONE
}

type GOAL struct {
	ID      string
	Name    string // One word descriptive title, e.g. finishTaskOne or finish_task_one
	Details string
	DueDate int
}

type MILESTONE struct {
	ID    string
	Name  string // One word descriptive title, e.g. finishTaskOne or finish_task_one
	Tasks []string
}

type TASK struct {
	ID      string
	Name    string // One word descriptive title, e.g. finishTaskOne or finish_task_one
	Details string
	DueDate int
}

func init() {
	pathBase := fmt.Sprintf("%s/.docket", homeDir())
	pathDocket := pathBase + "/docket.json"
	// if docket hasn't been run before, create create /home/{user}/.docket/docket.json
	// else read in docket from /home/{user}/.docket/docket.json
	if _, err := os.Stat(pathBase); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathBase, os.ModePerm)
		check(err)
		fmt.Println("Hi")
		f, err := os.Create(pathDocket)
		check(err)
		defer f.Close()
	} else {
		docketBytes, err := os.ReadFile(pathDocket)
		check(err)
		if len(docketBytes) > 0 {
			err = json.Unmarshal(docketBytes, &Docket)
			check(err)
		}
	}
}

// Flush the contents of the current docket to /home/{user}/.docket/docket.json
func (docket *DOCKET) flush() {
	docketElementBytes, err := json.Marshal(docket.Elements)
	check(err)
	fmt.Println("Hi")
	f, err := os.OpenFile(fmt.Sprintf("%s/.docket/docket.json", homeDir()), os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	check(err)
	defer f.Close()
	_, err = f.Write(docketElementBytes)
	check(err)
}

func (docket *DOCKET) createElement(name string) {
	docket.Elements = append(docket.Elements, ELEMENT{
		ID:   generateID(name),
		Name: name,
	})
}

// docket create goal
func (docket *DOCKET) createGoal(elementName, name string, dueDate int) error {
	var details string
	fmt.Printf("Enter details: ")
	fmt.Scanln(&details)
	// implement multiple choice promptui to add goals, milestones: https://liza.io/implementing-multiple-choice-selection-in-go-with-promptui/

	for i, x := range docket.Elements {
		if x.Name == elementName {
			docket.Elements[i].Goals = append(docket.Elements[i].Goals, GOAL{
				ID:      generateID(name),
				Name:    name,
				Details: details,
				DueDate: dueDate,
			})
		}
	}

	return nil
}

// docket create milestone
func (docket *DOCKET) createMilestone(elementName, name string) {
	// implement multiple choice promptui to add tasks, goals: https://liza.io/implementing-multiple-choice-selection-in-go-with-promptui/
}

// docket create task
func (docket *DOCKET) createTask(elementName, name string, dueDate int) error {
	var details string
	fmt.Printf("Enter details: ")
	fmt.Scanln(&details)
	// implement multiple choice promptui to add goals, milestones: https://liza.io/implementing-multiple-choice-selection-in-go-with-promptui/

	for i, x := range docket.Elements {
		if x.Name == elementName {
			docket.Elements[i].Tasks = append(docket.Elements[i].Tasks, TASK{
				ID:      generateID(name),
				Name:    name,
				Details: details,
				DueDate: dueDate,
			})
		}
	}

	return nil
}

// docket search
func (docket *DOCKET) search() {}

// docket glance
func (docket *DOCKET) glance(period string) {
	switch period {
	case "day":
		fmt.Println("Current Day.")
	case "week":
		fmt.Println("Current Week.")
	case "month":
		fmt.Println("Current Month.")
	case "year":
		fmt.Println("Current Year.")
	}
}

// Execute a command @ stdin, receive stdout
func execCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Error checking
func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Return current user's home directory as string
func homeDir() string {
	currentUser, err := user.Current()
	check(err)

	return fmt.Sprintf("/home/%s", currentUser.Username)
}

// Generate ID from struct name
func generateID(name string) string {
	newSha := sha256.New()
	_, err := newSha.Write([]byte(name + strconv.FormatInt(time.Now().Unix(), 10)))
	check(err)

	return string(fmt.Sprintf("%x", newSha.Sum(nil)))
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
