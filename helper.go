package main

import "strings"

func ValidateUserInput(firstNameRef string, lastNameRef string, emailRef string, userTicketsRef uint) (bool, bool, bool) {
	isValidName := len(firstNameRef) >= 2 && len(lastNameRef) >= 2
	isValidEmail := strings.Contains(emailRef, "@")
	isValidTicketNumber := userTicketsRef > 0 && userTicketsRef <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
