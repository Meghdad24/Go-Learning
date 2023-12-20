package main

import "fmt"

func main() {
	ConferenceName := "Go Conference"
	const ticketCounts int = 40
	var ticketSold int = 0
	var bookings [50]string

	fmt.Printf("Hello! Welcome to %v!\n", ConferenceName)
	fmt.Printf("We have %v Ticket for sell.\n", ticketCounts-ticketSold)

	var firstName, lastName string
	fmt.Println("What's your name(Firstname & lastname)?")
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	bookings[0] = firstName + " " + lastName

	var userEmail string
	fmt.Println("Enter your email address:")
	fmt.Scan(&userEmail)

	var userTickets int
	fmt.Println("How many tickets you want?")
	fmt.Scan(&userTickets)

	ticketSold += userTickets

	fmt.Printf("Mr.%v! You have %v tickets and %v tickets are left for sale.", lastName, userTickets, ticketCounts-ticketSold)

	fmt.Println("Array properties:")
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Size of array: %v", len(bookings))
}
