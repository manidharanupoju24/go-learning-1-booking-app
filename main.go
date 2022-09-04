package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	//remainingTickets := 50 // short hand notation doesn't work with const
	var remainingTickets uint = 50
	var bookings []string //string type slice , use append method to add data to slices
	//var bookings[50]string  //string type array

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("Hello world!")
	fmt.Printf("Welcome to %v, booking application\n", conferenceName) // placeholder
	fmt.Println("We have a total of", conferenceTickets, "tickets")
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("There are %v conference tickets available\n", remainingTickets)
	//fmt.Println(conferenceName)

	//bookings[0] = "Manidhar"
	//bookings[1] = "Nicole"

	for {
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
		fmt.Printf("The first names for bookings are %v\n", firstNames)

		// loops are simplified in go. There are no while loops, do-while loops or for-each loops
		// there is only one loop called for loop
	}
}
