package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Printf("Enter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter Your last Name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter Your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	// firstName = "Tom"
	userTickets = 5
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
