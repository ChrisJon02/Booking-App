package main

import (
	"fmt"
)

//Package Level Variables

const conferenceTickets = 50

var conferenceName = "Music Show"
var remainingTickets int = 50

// Declared a struct
var bookings = make([]UserData, 0)

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets int
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidFirstName, isValidLastName, isValidEmail, isValidTickets := validateInputData(firstName, lastName, email, userTickets)

		if isValidFirstName && isValidLastName && isValidEmail && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)

			//creates a new goroutine
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our %v is booked out. Come back next year", conferenceName)

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
