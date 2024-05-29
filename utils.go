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
	Calendar []CALENDAR
	Todo     []TODO
}

type CALENDAR struct {
	ID     string
	Events []EVENT
}

type TODO struct {
	ID    string
	Tasks []TASK
}

type EVENT struct {
	ID        string
	Name      string
	Details   string
	Timestamp int
	Duration  int
}

type TASK struct {
	ID      string
	Details string
	DueDate int
}

type TASKWITHSUBTASKS struct {
	ID      string
	Details string
	DueDate int
	Tasks   []TASK
}

func createCalendar(name string) {
	err := os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", homeDir(), name), []byte("c"), 0777)
	check(err)
}

func createTodo(name string) {
	err := os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", homeDir(), name), []byte("t"), 0777)
	check(err)
}

func renameCalendar(oldName, newName string) {}
func renameTodo(oldName, newName string)     {}

func parseCalendar() {}
func parseTodo()     {}

// Format: d add event {calendarName} {name} {timestamp} {duration}
func createEvent(calendarName, name string, timestamp, duration int) {
	var details string
	fmt.Printf("Enter event details: ")
	fmt.Scanln(&details)
	eventJson, err := json.Marshal(EVENT{
		ID:        generateID(name),
		Name:      name,
		Details:   details,
		Timestamp: timestamp,
		Duration:  duration,
	})
	check(err)

	err = os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", homeDir(), calendarName), eventJson, 0644)
	check(err)
}

// Format: d add task {calendarName} {name} {timestamp} {duration}
func createTask(calendarName, name string, dueDate int) {
	var details string
	fmt.Printf("Enter task details: ")
	fmt.Scanln(&details)
	taskJson, err := json.Marshal(TASK{
		ID:      generateID(name),
		Details: details,
		DueDate: dueDate,
	})
	check(err)

	err = os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", homeDir(), calendarName), taskJson, 0644)
	check(err)
}

// func CreateTaskWithSubTask() {}
func search() {}

func glance(period string) {
	if period == "day" {

	} else if period == "week" {

	} else if period == "month" {

	} else if period == "year" {

	}
}

func execCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func homeDir() string {
	currentUser, err := user.Current()
	check(err)

	return fmt.Sprintf("/home/%s", currentUser.Username)
}

func generateID(name string) string {
	newSha := sha1.New()
	_, err := newSha.Write([]byte(name + strconv.FormatInt(time.Now().Unix(), 10)))
	check(err)

	return string(newSha.Sum(nil))
}
