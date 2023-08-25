package helper

import "strings"

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (x, y, z bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
