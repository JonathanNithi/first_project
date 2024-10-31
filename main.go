package main

import "fmt"

func main() {
	//initialise variables and constants
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//welcome message
	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("=================================================")
	fmt.Println()
	fmt.Println("We have a total of", conferenceTickets, "tickets")
	fmt.Println("The amount of remaining tickets are", remainingTickets)
	fmt.Println()
	fmt.Println("Get the tickets to your favorite conference")
}
