package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings []string


func main() {
	
	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if  isValidName && isValidEmail && isValidTicketCount {	
			bookTicket(userTickets, firstName, lastName, email)
			
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
			
		} else {
			if !isValidName {
				fmt.Printf("First name or last name does not meet the required amount of characters\n")
			} 
			if !isValidEmail {
				fmt.Printf("Email address you entered is invalid, it must contain an @ symbol\n")
			} 
			if !isValidTicketCount {
				fmt.Printf("Please select a valid amount of tickets\n")
			}
			// will start a new iteration immediately
			continue
		} 

	}
	
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here for the event")
}

func getFirstNames() []string {
	firstNames := []string{}
	// _ is used to define unneeded variables -- we tell Go to explicitly not use
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput (firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets
	// we can return multiple values in Go
	return isValidName, isValidEmail, isValidTicketCount
}

func getUserInput () (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like to buy?")
	fmt.Scan(&userTickets)
	
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("%v %v booked %v tickets. An email verification will be sent to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}