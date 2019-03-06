package scheduler

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"../appointment"
)

type scheduler struct {
	name string
	isRunning bool
	appointments []appointment.Appointment
}

// Scheduler Methods 

func Init(name string) *scheduler {
	s := &scheduler {
		name: name,
		isRunning: true,
	}

	for s.IsRunning() {
		s.DisplayMenu()
	}

	return s
}

func (scheduler *scheduler) IsRunning() bool {
	return scheduler.isRunning
}

func (scheduler *scheduler) DisplayMenu() {
	fmt.Println("=====================")
	fmt.Println("Select one of the following options:")
	fmt.Println("1. List all future appointments")
	fmt.Println("2. Schedule a new appointment")
	fmt.Println("3. Exit application")
	fmt.Println("=====================")

	selection := scheduler.getInput("option")

	isDigit, _ := regexp.MatchString("[0-9]{1}", selection)

	if (isDigit) {
		s, _ := strconv.Atoi(selection)
		switch s {
			case 1:
				scheduler.displayAppointments()
			case 2:
				scheduler.addAppointment()
			case 3:
				scheduler.isRunning = false
		}
	} else {
		fmt.Println("Please enter a valid option (1-3): ")
	}
}

func (scheduler *scheduler) getInput(setting string) string {
	var value string

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	if setting == "option" {
		value = string(input[0])
	} else if setting == "string" {
		value = string(input)
	}

	return value
}

func (scheduler *scheduler) addAppointment() {
	newAppointment := *appointment.New()
	scheduler.appointments = append(scheduler.appointments, newAppointment)
}

func (scheduler *scheduler) displayAppointments() {
	Tracker := *appointment.Tracker
	timeslots := Tracker.TimeSlots

	for i := 0; i < len(timeslots); i++ {
		for j := 0; j < len(timeslots[i]); j++ {
			appointment := timeslots[i][j].Appointment;
			if appointment != nil {
				fmt.Println("------------------")
				fmt.Println("Customer: ", appointment.Customer)
				fmt.Println("Services: ", appointment.Service)
				fmt.Println("Time: ", appointment.Timeslot)
				fmt.Println("Duration: ", appointment.Duration)
				fmt.Println("------------------")
			}
		}
	}
}