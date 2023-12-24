package main

import "fmt"

func main() {
	pastNumber := 1
	currentNumber := 1
	var temp int
	for currentNumber <= 21 {
		for i := 0; i < currentNumber; i++ {
			fmt.Print(currentNumber)
		}
		fmt.Println("")
		temp = currentNumber
		currentNumber = currentNumber + pastNumber
		pastNumber = temp
	}

}
