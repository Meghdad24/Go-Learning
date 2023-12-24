package main

import (
	"fmt"
)

func main() {
	var myMap map[int]string = map[int]string{
		1: "One",
		2: "Tow",
	}
	// myMap := make(map[int]string)
	myKey := 13
	myMap[myKey] = "thirteen"
	fmt.Println(myMap)        //map[13:thirteen]
	fmt.Println(myMap[myKey]) //thirteen
}
