package main

import (
	"fmt"
)

// func main() {
// 	defer fmt.Println("world")
// 	fmt.Println("hello")
// }

func main() {
	defer func() { fmt.Println("In inline defer") }()
	fmt.Println("Executed")
}
