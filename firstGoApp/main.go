package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Gany-Con"
	const totalTickets = 50
	var ticketsRemaining = totalTickets

	fmt.Printf("Welcome to my %v booking app\n", conferenceName)
	fmt.Printf("We have %v tickets available\n", ticketsRemaining)

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var ticketsRequested int
		fmt.Printf("Please enter your first name\n")
		fmt.Scan(&firstName)
		fmt.Printf("Please enter your last name\n")
		fmt.Scan(&lastName)
		fmt.Printf("Please enter your email\n")
		fmt.Scan(&email)
		fmt.Printf("Please enter number of tickets\n")
		fmt.Scan(&ticketsRequested)

		if ticketsRemaining >= ticketsRequested {
			ticketsRemaining -= ticketsRequested

		}

		for _, booking := range bookings {
			var name = strings.Fields(booking)
			var firstName = name[0]
			fmt.Printf("This bookings first name %v", firstName)
		}

		var fullName = firstName + lastName

		bookings = append(bookings, fullName)

		fmt.Printf("You will recieve a confirmation email at %v for booking %v tickets\n", email, ticketsRequested)

		fmt.Printf("Current bookings %v", bookings)
	}

}
