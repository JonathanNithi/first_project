package main

func validateUserInput(firstName string, lastName string, ticketCount uint) (bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidTicket bool = ticketCount > 0 && ticketCount <= remainingTickets

	return isValidName, isValidTicket
}
