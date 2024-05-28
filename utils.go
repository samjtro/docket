package main

type DOCKET struct {
	Calendar []EVENT
	Todo []TASK
}

type Calendar struct {
	ID string
	Events []EVENT
}

type Todo struct {
	ID string
	Tasks []TASK
}

type EVENT struct {
	ID string
	Name string
	Details string
	Timestamp int
	Duration int
	Location string
}

type TASK struct {
	ID string
	Details string
	Timestamp int
}

func CreateCalendar() {
	
}

func CreateEvent() {
	
}

func CreateTask() {
	
}

func Search() {

}

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
