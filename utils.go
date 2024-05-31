package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"
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

// Flush the contents of the current docket to /home/{user}/.docket/docket.json
func (docket *DOCKET) flush() {
	docketElementBytes, err := json.Marshal(docket.Elements)
	check(err)
	f, err := os.OpenFile(fmt.Sprintf("%s/.docket/docket.json", homeDir()), os.O_TRUNC, 0755)
	check(err)
	defer f.Close()
	f.Write(docketElementBytes)
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
	newSha := sha1.New()
	_, err := newSha.Write([]byte(name + strconv.FormatInt(time.Now().Unix(), 10)))
	check(err)

	return string(newSha.Sum(nil))
}
