package main

import (
	"fmt"
	"strings"
)

//Package Level Variables

const conferenceTickets = 50

var conferenceName = "Music Show"
var remainingTickets int = 50

// Declared a slice instead of array which is dynamic in size
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidFirstName, isValidLastName, isValidEmail, isValidTickets := validateInputData(firstName, lastName, email, userTickets)

		if isValidFirstName && isValidLastName && isValidEmail && isValidTickets {

			bookings = bookTicket(userTickets, bookings, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our %v is booked out. Come back next year", conferenceName)
				break
			}
		} else {
			if !isValidFirstName {
				fmt.Println("Your First name is too short. Please enter valid name")
			}
			if !isValidLastName {
				fmt.Println("Your Last name is too short. Please enter valid name")
			}
			if !isValidEmail {
				fmt.Println("Your Email-Id is invalid")
			}
			if !isValidTickets {
				if userTickets <= 0 {
					fmt.Println("We can't book tickets less than 1")
				} else {
					if remainingTickets == 1 {
						fmt.Printf("We only have 1 more ticket remaining, we can't book %v tickets\n", userTickets)
					} else {
						fmt.Printf("We only have %v more tickets remaining, we can't book %v tickets\n", remainingTickets, userTickets)
					}
				}
			}

		}
	}

}

// function to print greet users
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Book your tickets here for the %v\n", conferenceName)
}

// function to get the first names of the bookings to maintain privacy of the users
func getFirstNames() []string {
	firstNames := []string{}

	//range for loop - a for loop to run through arrays, slices
	// "_" is the blank identifier used when you don't need a particular variable
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

// function to validate input data
func validateInputData(firstName string, lastName string, email string, userTickets int) (bool, bool, bool, bool) {

	isValidFirstName := len(firstName) >= 2
	isValidLastName := len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidFirstName, isValidLastName, isValidEmail, isValidTickets
}

// function to get user input
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

func bookTicket(userTickets int, bookings []string, firstName string, lastName string, email string) []string {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)

	return bookings
}
