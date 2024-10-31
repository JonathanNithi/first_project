package main

import "fmt"

func main() {
	//initialise variables and constants
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	//welcome message
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Println("=================================================")
	fmt.Println()
	fmt.Printf("We have a total of %v tickets\n", conferenceTickets)
	fmt.Printf("The amount of remaining tickets are %v\n", remainingTickets)
	fmt.Println()
	fmt.Println("Get the tickets to your favorite conference")

	//Collect user input
	var firstName string
	var lastName string
	var ticketCount uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter the count of tickets you wish to purchase:")
	fmt.Scan(&ticketCount)

	fmt.Println()
	fmt.Printf("User %v %v has purchased %v tickets.\n", firstName, lastName, ticketCount)
}
