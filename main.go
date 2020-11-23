package main

import (
	"./schedule"
)

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
