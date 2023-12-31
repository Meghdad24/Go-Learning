package main

import (
	"fmt"
)

// func main() {
// 	defer fmt.Println("world")
// 	fmt.Println("hello")
// }

// func main() {
// 	defer func() { fmt.Println("In inline defer") }()
// 	fmt.Println("Executed")
// }

func main() {
	i := 0
	i = 1
	defer fmt.Println(i)
	i = 2
	defer fmt.Println(i)
	i = 3
	defer fmt.Println(i)
}
