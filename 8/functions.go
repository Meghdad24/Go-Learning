package main

import "fmt"

//number 1
func power(num, power int) int {
	output := 1
	for i := 0; i < power; i++ {
		output = output * num
	}
	return output
}

//number 2
func sum(nums ...int) (total int) {
	total = 0

	fmt.Print(nums, " =sum=> ")

	for _, num := range nums {
		total += num
	}
	return
}

func main() {
	fmt.Println(power(5, 3))

	fmt.Println(sum(45, 23, 56, 45, 78, 23, 45, 11, 6, 23, 4, 856, -100))
}
