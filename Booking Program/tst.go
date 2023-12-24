package main

import "fmt"

func main() {
	ConferenceName := "Go Conference"
	const ticketCounts int = 40
	var ticketSold int = 0
	var bookingArray [40]string
	var bookingSlice []string

	fmt.Printf("Hello! Welcome to %v!\n", ConferenceName)
	fmt.Printf("We have %v Ticket for sell.\n", ticketCounts-ticketSold)

	var firstName, lastName string
	fmt.Println("What's your name(Firstname & lastname)?")
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	bookingArray[0] = firstName + " " + lastName
	bookingSlice = append(bookingSlice, firstName+" "+lastName)

	var userEmail string
	fmt.Println("Enter your email address:")
	fmt.Scan(&userEmail)

	var userTickets int
	fmt.Println("How many tickets you want?")
	fmt.Scan(&userTickets)

	ticketSold += userTickets

	fmt.Printf("Mr.%v! You have %v tickets and %v tickets are left for sale.", lastName, userTickets, ticketCounts-ticketSold)

	fmt.Println("Array properties:")
	fmt.Printf("The whole array: %v\n", bookingArray)
	fmt.Printf("The first value: %v\n", bookingArray[0])
	fmt.Printf("Array type: %T\n", bookingArray)
	fmt.Printf("Size of array: %v\n", len(bookingArray))

	fmt.Println("\nSlice properties:")
	fmt.Printf("The whole slice: %v\n", bookingSlice)
	fmt.Printf("The first value: %v\n", bookingSlice[0])
	fmt.Printf("Slice type: %T\n", bookingSlice)
	fmt.Printf("Size of slice: %v", len(bookingSlice))
}
