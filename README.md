# Hair Salon Scheduler

 This is a simple command-line application designed to list and schedule non-conflicting (start times) appointments for: Haircuts, Haircuts & Shampoo, or a custom order. It contains three main structs and two pacakges (Scheduler, Appointent, Tracker | Scheduler, Appointment) each with their own respective methods. It will only allow the hair stylist to book future appointments, but still keep a temporary log of the day's appointments (until exit).

## Getting Started

You'll need to have the latest version of GO installed to recompile the source code if you decide to make modifications.

### Prerequisites

Latest Version of GO
https://golang.org/README.md

### Installing

No "installation" needed. Unzip and run the standalone command-line application.

## Running the test

1. Unzip application
2. Run "app"
3. Choose "1" to view all appointments (will be null on start)
   Choose "2" to start the scheduling process (continue to step 4)
   Choose "3" to exit the application loop
4. Option "2":
   - Enter customer name
   - Enter date (00:00am,AM,pm,PM)
   - Choose service (1 = Haircut (30 mins),2 = Haircut & Shampoo (1hr),3 = Custom Order (30mins))
5. Return to Step 3 and Choose "1" to view newly added appointment
6. Choose "2" and try creating an appointment with the same time as before (prompts for another time)
7. Try entering something invalid: "trdfeygfh", "54:68pm", "13:45am", etc. (prompts for another time)
8. Feel free to create another appointment and then Choose "1" from the main menu

## Built With

* [Go](https://golang.org/doc/) - GO Programming Langauge

## Authors

* **Everett Morgan** - *Entire project* - [Website](https://ejmorgan.com/)

## Acknowledgments

* Go Tour!
