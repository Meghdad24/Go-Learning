package main

import "fmt"

func main() {
	ConferenceName := "Go Conference"
	const ticketCounts int = 40
	var ticketSold int = 0

	fmt.Printf("Hello! Welcome to %v!\n", ConferenceName)
	fmt.Printf("We have %v Ticket for sell.\n", ticketCounts-ticketSold)

}
