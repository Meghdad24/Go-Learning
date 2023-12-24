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

	x := make(map[string]int)
	a, ok := x["num"]
	fmt.Println(a, ok) // 0 false

	x["num"] = 0
	b, ok := x["num"]
	fmt.Println(b, ok, "\n")

	var companyProfile = map[string]string{
		"name":    "companyName",
		"address": "sampleAddress",
	}
	var editorMap = companyProfile // == editorMap := companyProfile

	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])
	//companyName 	 sampleAddress
	fmt.Println(editorMap["name"], "\t", editorMap["address"])
	//companyName 	 sampleAddress

	editorMap["name"] = "new name"
	editorMap["address"] = "new address"

	//reference map also edited when editor map edit
	fmt.Println(companyProfile["name"], "\t", companyProfile["address"])
	//new name 	 new address
	fmt.Println(editorMap["name"], "\t", editorMap["address"], "\n")
	//new name 	 new address

	animals := make(map[int][]string) // nil map of string-int pairs
	animals[0] = []string{"Gopher", "running", "rodent"}
	animals[1] = []string{"owl", "flying", "carnivorous"}
	animals[2] = []string{"cheetah", "running", "carnivorous"}
	animals[3] = []string{"eagle", "flying", "carnivorous"}
	animals[4] = []string{"lion", "running", "carnivorous"}

	for index, animal := range animals {
		fmt.Printf("%v- %s is %s animal and can %s \n", index, animal[0], animal[2], animal[1])
	}
}
