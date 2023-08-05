package main

import (
	"fmt"
	"strings"
	"time"
)

// welcome message for the users
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Book your tickets here for the %v\n", conferenceName)
}

// gets user input
func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// validates input data
func validateInputData(firstName string, lastName string, email string, userTickets int) (bool, bool, bool, bool) {

	isValidFirstName := len(firstName) >= 2
	isValidLastName := len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidFirstName, isValidLastName, isValidEmail, isValidTickets
}

// get first names of the users
func getFirstNames() []string {
	firstNames := []string{}

	//range for loop - a for loop to run through arrays, slices
	// "_" is the blank identifier used when you don't need a particular variable
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	return firstNames
}

// book tickets for the user
func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct for a user for storing all their details
	var userData = UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings : %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)
}

// Send ticket to the user's email
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("Booked %v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n%v \nTo %v\n", ticket, email)
	fmt.Println("################")
}
