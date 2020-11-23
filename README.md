# ScheduleCLI

 ScheduleCLI provides a simple api to create different types of schedules based on appointment type. (i.e. hair salon vs. tax accounting)
 
 ```go
import "github.com/everettmorgan/schedule-cli"
 
func main() {
	c := schedule.NewScheduleCfg("Thee Hair Salon", 8, 13, "PST")
	a := schedule.AppointmentTypes{
		schedule.NewAppointmentType("trim", 10, 15),
		schedule.NewAppointmentType("haircut", 20, 30),
		schedule.NewAppointmentType("shampoo", 10, 15),
		schedule.NewAppointmentType("custom", 30, 60),
	}

	s := schedule.NewSchedule(&c, &a)
	s.Init()
}
 ```

## Getting Started

### Prerequisites

Latest Version of GO: https://golang.org/README.md

## Authors

* **Everett Morgan** - *Entire project* - [Website](https://ejmorgan.com/)
