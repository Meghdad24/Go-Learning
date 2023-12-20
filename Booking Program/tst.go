package main

import "fmt"

func main() {
	ConferenceName := "Go Conference"
	const ticketCounts int = 40
	var ticketSold int = 0

	fmt.Printf("Hello! Welcome to %v!\n", ConferenceName)
	fmt.Printf("We have %v Ticket for sell.\n", ticketCounts-ticketSold)

	var firstName, lastName string
	fmt.Println("What's your name(Firstname & lastname)?")
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)

	var userEmail string
	fmt.Println("Enter your email address:")
	fmt.Scan(&userEmail)

	var userTickets int
	fmt.Println("How many tickets you want?")
	fmt.Scan(&userTickets)

	ticketSold += userTickets

	fmt.Printf("Mr.%v! You have %v tickets and %v tickets are left for sale.", lastName, userTickets, ticketCounts-ticketSold)

}
