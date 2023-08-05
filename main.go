package main

import (
	"fmt"
	"strconv"
)

//Package Level Variables

const conferenceTickets = 50

var conferenceName = "Music Show"
var remainingTickets int = 50

// Declared an empty List of maps unnlike a single map
var bookings = make([]map[string]string, 0)

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
		firstNames = append(firstNames, booking["First name"])
	}
	return firstNames

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

func bookTicket(userTickets int, bookings []map[string]string, firstName string, lastName string, email string) []map[string]string {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user storing all their details
	var userData = make(map[string]string)
	userData["First name"] = firstName
	userData["Last name"] = lastName
	userData["Email"] = email
	userData["Number of Tickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings : %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)

	return bookings
}
