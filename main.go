package main

import (
	"fmt"
	"strings"
)

func main() {
	//initialise variables and constants
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookingsList = []string{}

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

	for {
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter the count of tickets you wish to purchase:")
		fmt.Scan(&ticketCount)

		if ticketCount > remainingTickets {
			fmt.Printf("The remaining ticket count is %v and therefore, you cannot purchase %v tickets\n", remainingTickets, ticketCount)
			continue
		}

		fmt.Println()
		fmt.Printf("User %v %v has purchased %v tickets.\n", firstName, lastName, ticketCount)
		bookingsList = append(bookingsList, firstName+" "+lastName)

		remainingTickets = remainingTickets - ticketCount

		fmt.Printf("There are %v tickets left for %v\n", remainingTickets, conferenceName)
		fmt.Println("Hurry up and get your tickets before they get sold out!")

		firstNameList := []string{}
		//for loop with an unused variable instead of index
		for _, booking := range bookingsList {
			var names = strings.Fields(booking)
			firstNameList = append(firstNameList, names[0])
		}
		fmt.Printf("This is the list of bookings: %v \n", firstNameList)

		if remainingTickets == 0 {
			fmt.Printf("%v tickets are now sold out!\n", conferenceName)
		}
	}

}
