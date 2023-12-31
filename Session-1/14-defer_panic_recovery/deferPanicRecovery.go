package main

import (
	"fmt"
	"runtime/debug"
)

//----------------------------v-DEFER--------------------------
// func main() {
// 	defer fmt.Println("world")
// 	fmt.Println("hello")
// }

// func main() {
// 	defer func() { fmt.Println("In inline defer") }()
// 	fmt.Println("Executed")
// }

// func main() {
// 	i := 0
// 	i = 1
// 	defer fmt.Println(i)
// 	i = 2
// 	defer fmt.Println(i)
// 	i = 3
// 	defer fmt.Println(i)
// }

// func main() {
// 	i := 1
// 	defer fmt.Println(i)
// 	i++
// 	fmt.Println(i)
// 	fmt.Println("First")
// }

//---------------------------PANIC------------------------------

// func main() {

// 	a := []string{"a", "b"}
// 	print(a, 2)
// }

// func print(a []string, index int) {
// 	fmt.Println(a[index])
// }

// func main() {

// 	a := []string{"a", "b"}
// 	checkAndPrint(a, 2)
// }

// func checkAndPrint(a []string, index int) {
// 	if index > (len(a) - 1) {
// 		panic("Out of bound access for slice")
// 	}
// 	fmt.Println(a[index])
// }

// func main() {

// 	a := []string{"a", "b"}
// 	checkAndPrint(a, 2)
// 	fmt.Println("Exiting normally")
// }

// func checkAndPrint(a []string, index int) {
// 	defer handleOutOfBounds()
// 	if index > (len(a) - 1) {
// 		panic("Out of bound access for slice")
// 	}
// 	fmt.Println(a[index])
// }

// func handleOutOfBounds() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovering from panic:", r)
// 	}
// }

func main() {
	a := []string{"a", "b"}
	checkAndPrint(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrint(a []string, index int) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
		fmt.Println("Stack Trace:")
		debug.PrintStack()
	}
}
