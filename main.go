package main

import (
	"fmt"
	"strings"
)

// initialise variables and constants
const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookingsList = []string{}

func main() {

	//welcome message
	greetUsers()

	for {
		//Collect user input
		firstName, lastName, ticketCount := getUserInput()

		//validate user input
		isValidName, isValidTicket := validateUserInput(firstName, lastName, ticketCount)

		if isValidName && isValidTicket {

			//book tickets
			bookingsList = bookTickets(firstName, lastName, ticketCount)

			//print first name of who booked the tickets
			printFirstNames(bookingsList)

			if remainingTickets == 0 {
				fmt.Printf("%v tickets are now sold out!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a name which is longer than 2 letters")
			}
			if !isValidTicket {
				if ticketCount <= 0 {
					fmt.Println("You cannot enter a zero or a negative number for the required number of tickets")
				} else {
					fmt.Printf("The remaining ticket count is %v and therefore, you cannot purchase %v tickets\n", remainingTickets, ticketCount)
				}
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Println("=================================================")
	fmt.Println()
	fmt.Printf("We have a total of %v tickets\n", conferenceTickets)
	fmt.Printf("The amount of remaining tickets are %v\n", remainingTickets)
	fmt.Println()
	fmt.Println("Get the tickets to your favorite conference")
}

func printFirstNames(bookingsList []string) {
	firstNameList := []string{}
	//for loop with an unused variable instead of index
	for _, booking := range bookingsList {
		var names = strings.Fields(booking)
		firstNameList = append(firstNameList, names[0])
	}
	fmt.Printf("This is the list of bookings: %v \n", firstNameList)
}

func validateUserInput(firstName string, lastName string, ticketCount uint) (bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidTicket bool = ticketCount > 0 && ticketCount <= remainingTickets

	return isValidName, isValidTicket
}

func getUserInput() (string, string, uint) {

	var firstName string
	var lastName string
	var ticketCount uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter the count of tickets you wish to purchase:")
	fmt.Scan(&ticketCount)

	return firstName, lastName, ticketCount

}

func bookTickets(firstName string, lastName string, ticketCount uint) []string {
	fmt.Println()
	fmt.Printf("User %v %v has purchased %v tickets.\n", firstName, lastName, ticketCount)
	bookingsList = append(bookingsList, firstName+" "+lastName)

	remainingTickets = remainingTickets - ticketCount

	fmt.Printf("There are %v tickets left for %v\n", remainingTickets, conferenceName)
	fmt.Println("Hurry up and get your tickets before they get sold out!")

	return bookingsList
}
