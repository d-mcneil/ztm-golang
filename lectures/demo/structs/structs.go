package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	IsBoarded    bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{Name: "Casey", TicketNumber: 123, IsBoarded: false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 456}
		ella = Passenger{Name: "Ella", TicketNumber: 789}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 101
	fmt.Println(heidi)

	casey.IsBoarded = true
	bill.IsBoarded = true

	if bill.IsBoarded {
		fmt.Println(bill.Name, "has boarded the bus.")
	}

	if casey.IsBoarded {
		fmt.Println(casey.Name, "has boarded the bus.")
	}

	heidi.IsBoarded = true
	bus := Bus{FrontSeat: heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is sitting in the front seat.")

}
