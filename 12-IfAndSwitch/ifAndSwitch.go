package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3
	if a > b {
		if a > c {
			fmt.Println("Biggest is a")
		} else if b > c {
			fmt.Println("Biggest is b")
		} else {
			fmt.Println("Biggest is c")
		}
	} else if b > c {
		fmt.Println("Biggest is b")
	} else {
		fmt.Println("Biggest is c")
	}

	//New MoDeL!
	if a := 6; a > 5 {
		fmt.Println("a is greater than 5")
	}
}
