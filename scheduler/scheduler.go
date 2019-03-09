package scheduler

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"../appointment"
)

// Scheduler manages the program loop and has a pointer to the tracker
type Scheduler struct {
	name      string
	isRunning bool
	calendar  string
	tracker   string
}

// New creates a scheduler and returns a pointer
func New(agentName string) *Scheduler {
	return &Scheduler{
		name:      agentName,
		isRunning: true,
	}
}

func (scheduler *Scheduler) displayMenuAndAwaitSelection() {
	displayText := `=====================
		Welcome to the Hair Salon Scheduler!
		What would you like to do?
		1. View Upcoming Appointments
		2. Schedule New Appointment
		3. Exit Application
		=====================`

	_, option := scheduler.displayThenGetInputOrSelection("option", displayText)

	// TODO: add handleOptionAction method
	scheduler.handleOptionAction(option)
}

func (scheduler *Scheduler) displayThenGetInputOrSelection(setting string, displayToUser string, options ...string) (valid bool, input string) {
	var rValid bool
	var rValue string

	fmt.Println(displayToUser)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	if setting == "option" {
		rValue = userInput
		rValid = true

		for i := 0; i < len(options); i++ {
			fmt.Println(string(i+1) + options[i])
		}

		result, _ := regexp.MatchString("^[0-9]$", rValue)

		if !result {
			rValid = false
		}

	} else if setting == "string" {
		rValue = string(userInput)
		rValid = false
	}

	return rValid, rValue
}

func (scheduler *Scheduler) addAppointment() {
	questions := map[string]string{
		"customer": "What's the customer's name?",
		"service":  "What kind of service is this?",
		"timeslot": "What time is the appointment? Please use the following format: 12:00AM/PM",
		"duration": "How long is this appointment?",
	}

	_, customer := scheduler.displayThenGetInputOrSelection("string", questions["customer"])
	_, service := scheduler.displayThenGetInputOrSelection("option", questions["service"], "Haircut", "Shampoo", "Custom")
	_, timeslot := scheduler.displayThenGetInputOrSelection("string", questions["timeslot"])
	_, duration := scheduler.displayThenGetInputOrSelection("option", questions["duration"], "15", "30", "45", "60")

	newAppointment := *appointment.New(customer, service, timeslot, duration)
	scheduler.appointments = append(scheduler.appointments, newAppointment)
}
