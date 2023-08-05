package main

import "strings"

// function to validate input data
func validateInputData(firstName string, lastName string, email string, userTickets int) (bool, bool, bool, bool) {

	isValidFirstName := len(firstName) >= 2
	isValidLastName := len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidFirstName, isValidLastName, isValidEmail, isValidTickets
}
