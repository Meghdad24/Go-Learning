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

	//Switch concept!

	switch ch := 'c'; ch {
	case 'a':
		fmt.Println("a")
	case 'b', 'c':
		fmt.Println("b or c")
		fallthrough
	case 'f':
		fmt.Println("f")
	}

	//New MoDeLL!
	i := 50
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
	case i < 100:
		fmt.Println("i is less than 50")
		fallthrough
	case false:
		fmt.Println("FALSE!!!")
	}
}
