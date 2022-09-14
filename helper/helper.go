package helper

import (
	"fmt"
	"strings"
)

// by changing the name of functions by capitalizing it's first letter, we are making it ready for export
// this is a unique functionality of go

func UserInputValidation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //pointers in go lang are called as special variables

	// pointers point to address of the variable in memory
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
