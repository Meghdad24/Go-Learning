package main

import "fmt"

func power(num int, power int) int {
	output := 1
	for i := 0; i < power; i++ {
		output = output * num
	}
	return output
}

func main() {
	m := 5
	p := 3
	fmt.Println(power(m, p))
}
