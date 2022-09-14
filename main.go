package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

// package level variables
//defined outside all the functions

var conferenceName string = "Go Conference"

const conferenceTickets = 50

//remainingTickets := 50 // short hand notation doesn't work with const
var remainingTickets uint = 50
var bookings []string //string type slice , use append method to add data to slices
//var bookings[50]string  //string type array

func main() {
	GreetUsers()

	//fmt.Println(conferenceName)
	//bookings[0] = "Manidhar"
	//bookings[1] = "Nicole"

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.UserInputValidation(firstName, lastName, email, userTickets, remainingTickets)
		/*
			if userTickets > remainingTickets {
				fmt.Printf("We only have %v tickets remaining, you can't book %v tickets.\n", remainingTickets, userTickets)
				continue
				// continue causes loop to skip the remainder of the body of the loop
				// immediately restest the condition (in our infinite loop, condition is always true)
			}
		*/
		if isValidName && isValidEmail && isValidTicketNumber {

			BookTicket(userTickets, firstName, lastName, email)
			firstNames := GetFirstNames()
			fmt.Printf("The first names for bookings are %v\n", firstNames)

			// loops are simplified in go. There are no while loops, do-while loops or for-each loops
			// there is only one loop called for loop

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Printf("Our Go Conference is booked out! Come back next year")
				break
			}
		} else {
			fmt.Println("Your input data is invalid. Please check below and try again :")
			if !isValidName {
				fmt.Println("Please check your first name and last name")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}
		/*else if userTickets == remainingTickets {
			fmt.Printf("You have booked all the remaining %v tickets available\n", remainingTickets)
			fmt.Println("After this transaction, no tickets will be available")
			break
		} else {
			fmt.Printf("We only have %v tickets remaining, you can't book %v tickets.\n", remainingTickets, userTickets)
			continue
		}*/
	}
	/*
		// Switch example
		city := "London"

		switch city {
		case "New York":
			//execute code for booking New York Conference tickets
		case "Singapore":
			//execute code for booking Singapore Conference tickets
		case "London":
			// some code here
		case "Berlin":
			// some code here
		case "Mexico city":
			// some code here
		case "Hong Kong", "New Delhi":
			// some code here
		default:
			fmt.Println("No valid city selected")
		}
	*/
}

func BookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	//bookings[49] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	/*
		//fmt.Printf("The whole array : %v\n", bookings)
		fmt.Printf("The whole slice : %v\n", bookings)
		fmt.Printf("The first value : %v\n", bookings[0])
		fmt.Printf("Array type : %T\n", bookings)
		fmt.Printf("Size of the array : %v\n", len(bookings))
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will shortly get a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func GreetUsers() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("Hello world!")
	fmt.Printf("Welcome to %v, booking application\n", conferenceName) // placeholder
	fmt.Println("We have a total of", conferenceTickets, "tickets")
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("There are %v conference tickets available\n", remainingTickets)
}

func GetFirstNames() []string {
	firstNames := []string{}
	//for each loop
	/*
		//here index should be given but we are not using it => give blank identifier _
		for index, booking := range bookings {
			var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with split elements
			firstNames = append(firstNames, names[0])
		}
	*/
	// blank identifier -> to ignore a varible you don't want to use
	// with go, you need to make unused variables explicit
	for _, booking := range bookings {
		var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with split elements
		firstNames = append(firstNames, names[0])
	}
	//range iterates over elements for different data structures (not only for arrays and slices)
	// for arrays and slices, range provides the index and value for each element
	//fmt.Printf("These are all our bookings: %v\n", bookings)
	return firstNames
}
