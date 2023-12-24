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
	fmt.Println(myMap)               //map[13:thirteen]
	fmt.Println(myMap[myKey] + "\n") //thirteen

	myMap2 := map[string]map[int]bool{
		"Sunday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
		"Monday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
		"Tuesday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
		"Wednesday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
		"Thursday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
		"Friday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
		"Saturday": {
			1: true,
			2: true,
			3: true,
			4: true,
		},
	}
	for k, _ := range myMap2 {
		fmt.Printf("%s ->	%v\n", k, myMap2[k])
	}
	fmt.Println()

}
