package main

import "fmt"

func loop() {
	var sum int

	for i := 0; i < 10; i++ {
		sum += i
	}

	println(sum)

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// //infinit loop
	// infinit := 0
	// for {
	// 	infinit++
	// }
	// fmt.Println("This line never executed!")

	//----------------------------For range-------------------------------

	letters := []string{"a", "b", "c"}

	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	letters2 := []string{"X", "Y", "Z"}

	fmt.Println("\nOnly letter")
	for i := range letters2 {
		fmt.Printf("letter: %s\n", letters2[i])
	}

	//--------------------------------Map Concept----------------------------------------

	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	//Iterating over all keys and values
	fmt.Println("Both Key and Value")
	for k, v := range sample {
		fmt.Printf("key: %s | value: %s\n", k, v)
	}

	//Iterating over only keys
	fmt.Println("\nOnly keys")
	for k := range sample {
		fmt.Printf("key: %s\n", k)
	}

	//Iterating over only values
	fmt.Println("\nOnly values")
	for _, v := range sample {
		fmt.Printf("value :%s\n", v)
	}

	//----------------------------------String---------------------------

	samplee := "a£b"

	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range samplee {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range samplee {
		fmt.Printf("Value:%s\n", string(letter))
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range samplee {
		fmt.Printf("Start Index: %d\n", i)
	}

	//------------------------------break & lable------------------------

	fmt.Println("\nBreak concept")
	summ := 0
	for {
		summ++
		if summ == 10 {
			break
		}
	}

	fmt.Printf("\n%v\nnow this line will execute\n\n", summ)

	letterss := []string{"a", "b", "c"}

	for i := 1; i < 10; i++ {
		// define a lable with name 'second' for this loop
	second:
		for i := 2; i < 9; i++ {
			for _, l := range letterss {
				if l == "b" {
					// break the loop with second lable
					break second
				}
				fmt.Println(l)
			}
		}
	}

	fmt.Print("\nContineu concpet\n")

	//--------------------------------continue--------------------------
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
