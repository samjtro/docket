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
	Calendar []EVENT
	Todo     []TASK
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

func CreateCalendar(name string) {
	err := os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", HomeDir(), name), []byte(""), 0777)
	Check(err)
}

// Format: d add event {calendarName} {name} {timestamp} {duration}
func CreateEvent(calendarName, name string, timestamp, duration int) {
	var details string
	fmt.Printf("Enter event details: ")
	fmt.Scanln(&details)
	eventJson, err := json.Marshal(EVENT{
		ID:        GenerateID(name),
		Name:      name,
		Details:   details,
		Timestamp: timestamp,
		Duration:  duration,
	})
	Check(err)

	err = os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", HomeDir(), calendarName), eventJson, 0644)
	Check(err)
}

// Format: d add task {calendarName} {name} {timestamp} {duration}
func CreateTask(calendarName, name string, dueDate int) {
	var details string
	fmt.Printf("Enter task details: ")
	fmt.Scanln(&details)
	taskJson, err := json.Marshal(TASK{
		ID:      GenerateID(name),
		Details: details,
		DueDate: dueDate,
	})
	Check(err)

	err = os.WriteFile(fmt.Sprintf("%s/.foo/docket/elements/%s.json", HomeDir(), calendarName), taskJson, 0644)
	Check(err)
}

// func CreateTaskWithSubTask() {}
func Search() {}

func Glance(period string) {
	if period == "day" {

	} else if period == "week" {

	} else if period == "month" {

	} else if period == "year" {

	}
}

func ExecCommand(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func Check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func HomeDir() string {
	currentUser, err := user.Current()
	Check(err)

	return fmt.Sprintf("/home/%s", currentUser.Username)
}

func GenerateID(name string) string {
	newSha := sha1.New()
	_, err := newSha.Write([]byte(name + strconv.FormatInt(time.Now().Unix(), 10)))
	Check(err)

	return string(newSha.Sum(nil))
}
