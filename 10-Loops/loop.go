package main

import "fmt"

func main() {
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
}
