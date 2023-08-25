package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := printFirstNames(bookings)
			fmt.Printf("These are all our bookings %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked. Come back next year.")
				break
			}

		} else {
			fmt.Println()

			if !isValidEmail {
				fmt.Println("email address you entered is invalid")
			}

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is inValid")
			}

			fmt.Println()

			// fmt.Println("Your Input Data is Invalid , Pls try again")
			// fmt.Printf("We have only %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}

}

func greetUsers(conferenceName string, conferenceTickets, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func printFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (x, y, z bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter Your last Name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter Your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets, userTickets uint, bookings []string, firstName, lastName, email, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
