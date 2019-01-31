package appointment

import (
	"fmt"
	"reflect"
	"bufio"
	"os"
	"strconv"
	"regexp"
	"strings"
	"time"
)

// Appointment

type Appointment struct {
	Customer string
	Service string
	Timeslot string
	Duration int
	Tracker *tracker
}

func New() *Appointment {
	// If: first appointment -> Then: create Tracker key
	if Tracker.Key == nil {
		createTrackerKey()
	}

	a := &Appointment {
		Tracker: Tracker,
	}
	
	a.setName()
	a.setTimeSlot()
	a.setService()

	return a
}

// Appointment Methods

func (appointment *Appointment) bookAppointment(apptTime string) {
	hour := Tracker.convertTimeToKey(apptTime)["hour"]
	minute := Tracker.convertTimeToKey(apptTime)["minute"]

	timeSlot := &Tracker.TimeSlots[hour][minute]

	if timeSlot.IsReserved != true && appointment.isFutureTime(apptTime){
		timeSlot.IsReserved = true
		timeSlot.Appointment = appointment
		appointment.Timeslot = apptTime
	} else {
		appointment.promptUser("Sorry, that time is either taken or in the past.")
		appointment.setTimeSlot()
	}
}

func (appointment *Appointment) isFutureTime(apptTime string) bool {
	currentHour, currentMin, _ := time.Now().Clock()
	currentTime := map[string]int {
		"hour": currentHour,
		"minute": currentMin,
	}

	if currentMin <= 00 {
		currentTime["minute"] = 0
	} else if (currentMin < 15) {
		currentTime["minute"] = 1
	} else if (currentMin > 15 && currentMin < 30) {
		currentTime["minute"] = 2
	} else if (currentMin > 30 && currentMin < 45) {
		currentTime["minute"] = 3
	}

	apptTimeParsed := Tracker.convertTimeToKey(apptTime)

	if apptTimeParsed["hour"] > currentTime["hour"] {
			return true
	} else if (apptTimeParsed["hour"] == currentTime["hour"]) {
		if apptTimeParsed["minute"] >= currentTime["minute"] {
			return true
		}
	}

	return false
}

func (appointment *Appointment) setName() {
	appointment.promptUser("What's the customer's name?")
	appointment.Customer = appointment.getInput("string")
}

func (appointment *Appointment) setTimeSlot() {
	appointment.promptUser("What time does it start? (ex. 01:15|30|45AM/PM)")
	timeSlot := appointment.getInput("string")
	validFormat, _ := regexp.MatchString("(12|01|02|03|04|05|06|07|08|09|10|11):(15|30|45|00)(am|pm|AM|PM)", timeSlot)

	if validFormat {
		appointment.bookAppointment(timeSlot)
	} else {
		appointment.promptUser("Please enter a valid time following the example format.")
		appointment.setTimeSlot()
	}
}

func (appointment *Appointment) setService() {
	appointment.promptUser("What kind of appointment", "Haircut", "Haircut & Shampoo", "Custom")
	selection, _ := strconv.Atoi(appointment.getInput("option"))

	isInt := reflect.TypeOf(selection).String()

	if (isInt == "int") {
		switch selection {
			case 0:
				appointment.Service = "Haircut"
				appointment.Duration = 30
			case 1:
				appointment.Service = "Haircut & Shampoo"
				appointment.Duration = 60
			case 2:
				appointment.promptUser("What's the custom order?")
				order := appointment.getInput("string")
				appointment.Service = order
				appointment.Duration = 30
		}

	} else {
		appointment.promptUser("Please enter a valid option (1-3)")
		appointment.setService()
	}
}

func (appointment *Appointment) promptUser(question string, options ...string) {
	fmt.Println(question)
	for i := 0; i < len(options); i++ {
		fmt.Println(i, ".", options[i])
	}
}

func (appointment *Appointment) getInput(setting string) string {
	var value string

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	switch setting {
		case "option":
			value = string(input[0])
		case "string":
			value = string(input)
	}

	return strings.Trim(value, "\n")
}

// Appointment Tracker

var Tracker = initTracker()

type tracker struct {
	TimeSlots [24][4]timeSlot
	Key map[string]map[string]int
}

type timeSlot struct {
	time string
	IsReserved bool
	Appointment *Appointment
}

func initTracker() *tracker {
	// Business Hour Options 
	const businessHoursStart int = 0
	const businessHoursEnd int = 23
	businessMinutes := []int{00, 15, 30, 45}

	// Create timetable matrix and populate with 24hr time
	// Selecting times -> matrix[0-23][0-3] = matrix[24hrs][00|15|30|45 mins]
	// ex. 4:30pm = matrix[16][2]
	matrix := [24][4]timeSlot{}

	for i, x := businessHoursStart, 0; i <= businessHoursEnd; i++ {
		for j, y := 0, 0; j < len(businessMinutes); j++ {
			var hour string
			var minutes string

			if i < 10 {
				hour = "0" + strconv.Itoa(i)
			} else {
				hour = strconv.Itoa(i)
			}

			if businessMinutes[j] == 00 {
				minutes = "00"
			} else {
				minutes = strconv.Itoa(businessMinutes[j])
			}

			t := timeSlot{
				time: (hour + minutes),
				IsReserved: false,
			}

			matrix[x][y] = t

			y++
		}

		x++
	}

	// Return Tracker pointer after timetable creation
	return &tracker{
		TimeSlots: matrix,
	}
}

func (tracker *tracker) convertTimeToKey(time string) map[string]int {
	// Convert string input to matrix pointers
	parsed := strings.Split(time, ":")
	parsedHour := parsed[0] + strings.ToUpper(parsed[1][2:4])
	parsedMinute := parsed[1][0:2]

	keyHour := Tracker.Key["hour"][parsedHour]
	keyMinute := Tracker.Key["minute"][parsedMinute]

	// Return object containing matrix pointers
	return map[string]int {
		"hour": keyHour,
		"minute": keyMinute,
	}
}

func createTrackerKey() {
	var hourKey = map[string]int {
		"12AM": 0,"01AM": 1, "02AM": 2, "03AM": 3,
		"04AM": 4,"05AM": 5, "06AM": 6, "07AM": 7,
		"08AM": 8, "09AM": 9, "10AM": 10, "11AM": 11,
		"12PM": 12, "01PM": 13, "02PM": 14, "03PM": 15,
		"04PM": 16, "05PM": 17, "06PM": 18, "07PM": 19,
		"08PM": 20, "09PM": 21, "10PM": 22, "11PM": 23,
	}

	var minuteKey = map[string]int {
		"00": 0, "15": 1,
		"30": 2, "45": 3,
	}

	Tracker.Key = map[string]map[string]int {
		"hour": hourKey,
		"minute": minuteKey,
	}
}