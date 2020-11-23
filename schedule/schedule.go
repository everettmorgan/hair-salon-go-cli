package schedule

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type option struct {
	desc string
	flag string
	vari []string
}

var menuopts = []option{
	{"view today's appointments", "t", []string{"today", "td", "now"}},
	{"view previous appointments", "p", []string{"past", "pt", "prev", "previous", "pr"}},
	{"view upcoming appointments", "f", []string{"future", "ft", "next", "upcoming"}},
	{"book an appointment", "b", []string{"book", "bk", "schedule"}},
	{"edit an appointment", "e", []string{"edit", "ed", "modify"}},
	{"delete an appointment", "d", []string{"delete", "del", "rm"}},
	{"save appointments to disk", "s", []string{"save", "sv", "wq"}},
	{"exit", "q", []string{"exit", "ex", "x"}},
}

var numopts = len(menuopts)
var menumap map[string]int

func createMenuMap() {
	menumap = make(map[string]int, numopts*2)
	for i, v := range menuopts {
		menumap[v.flag] = i
		for _, w := range v.vari {
			menumap[w] = i
		}
	}
}

var atypmap map[string]int

func createApptTypeMap() {
	atypmap = make(map[string]int, numopts*2)
	for i, v := range menuopts {
		for _, w := range v.vari {
			atypmap[w] = i
		}
	}
}

type appointment struct {
	customer string
	typ      *appointmentType
	start    time.Time
	end      time.Time
}

func (a *appointment) setName() error {
	name, err := ask("Enter the customer's name: ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.customer = name
	return nil
}

func (a *appointment) setType(types *AppointmentTypes) error {
	fmt.Printf("=== Appointment Types ===\n%s", types)
	apptt, err := ask("Choose the type of appointment: ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	ttype := strings.TrimRight(apptt, "\r\n")
	typidx, err := strconv.Atoi(ttype)
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.typ = &(*types)[typidx]
	return nil
}

func (a *appointment) setDate() error {
	startd, err := ask("Enter the appointment date (MM/DD/YY): ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	pstartd := strings.SplitN(startd, "/", -1)
	month, err := strconv.Atoi(pstartd[0])
	day, err := strconv.Atoi(pstartd[1])
	year, err := strconv.Atoi(pstartd[2])
	if err != nil {
		fmt.Println("invalid date")
		return err
	}

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		return err
	}

	apptT, err := ask("Enter appointment time (24:00): ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	papptt := strings.SplitN(apptT, ":", -1)
	appth, err := strconv.Atoi(papptt[0])
	apptm, err := strconv.Atoi(papptt[0])
	if err != nil {
		fmt.Println("invalid time")
		return err
	}

	a.start = time.Date(year, time.Month(month), day, appth, apptm, 0, 0, loc)
	return nil
}

// Schedule :
type Schedule struct {
	name  string
	appts []appointment
	types *AppointmentTypes
}

// NewSchedule :
func NewSchedule(cfg *scheduleCfg, atyps *AppointmentTypes) *Schedule {
	return &Schedule{
		name:  cfg.name,
		types: atyps,
	}
}

// Init :
func (s *Schedule) Init() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	createMenuMap()
	// createApptTypeMap()
	for {
		s.Menu()
	}
}

// Menu :
func (s *Schedule) Menu() {
	fmt.Printf("--- [MENU] %s ---\n", s.name)
	for i := 0; i < numopts; i++ {
		fmt.Printf("[%s] %s\n", menuopts[i].flag, menuopts[i].desc)
	}

	opt, err := ask("Please select an option: ")
	if err != nil {
		log.Fatal(err)
	}

	tsel := strings.TrimRight(opt, "\r\n")
	if tsel != menuopts[0].flag && menumap[tsel] == 0 {
		s.handle(&option{flag: tsel})
		return
	}

	s.handle(&menuopts[menumap[tsel]])
}

func (s *Schedule) handle(opt *option) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	cmd.Run()

	flag := strings.ToLower(opt.flag)

	switch flag {
	case "t":
		s.today()
		break
	case "p":
		s.past()
		break
	case "f":
		s.future()
		break
	case "b":
		s.book()
		break
	case "e":
		s.edit()
		break
	case "d":
		s.delete()
		break
	case "q":
		s.exit()
		break
	case "s":
		s.save()
		break
	default:
		s.invalid(flag)
		break
	}

	cmd.Run()
}

func (s *Schedule) today() {}

func (s *Schedule) past() {}

func (s *Schedule) future() {
	for i, v := range s.appts {
		fmt.Printf("[%d] {%s,%s} (%s)\t", i, v.customer, v.typ.name, v.start)
		if i%3 == 0 {
			fmt.Printf("\n")
		}
	}
}

func (s *Schedule) book() error {
	var a appointment
	err := a.setName()
	err = a.setType(s.types)
	err = a.setDate()
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.appts = append(s.appts, a)
	return nil
}

func (s *Schedule) edit() {}

func (s *Schedule) delete() {}

func (s *Schedule) save() {}

func (s *Schedule) exit() {
	os.Exit(1)
}

func (s *Schedule) invalid(flag string) {
	fmt.Printf("[%s] (ERROR) invalid option selected '%s'\n", time.Now().UTC(), flag)
}

type scheduleCfg struct {
	name  string
	start int
	end   int
	tz    string
}

// NewScheduleCfg :
func NewScheduleCfg(name string, start int, end int, tz string) scheduleCfg {
	return scheduleCfg{name, start, end, tz}
}

// AppointmentTypes :
type AppointmentTypes []appointmentType

// NewAppointmentType :
func NewAppointmentType(name string, price int, durmin int) appointmentType {
	return appointmentType{name, price, durmin}
}

func (at *AppointmentTypes) String() string {
	var str string
	for i, v := range *at {
		str += fmt.Sprintf("[%d] %s\n", i, v.name)
	}
	return str
}

type appointmentType struct {
	name   string
	price  int
	durmin int
}

func ask(message string) (string, error) {
	fmt.Print(message)
	r := bufio.NewReader(os.Stdin)
	inp, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	pinp := strings.TrimRight(inp, "\r\n")
	return pinp, nil
}
