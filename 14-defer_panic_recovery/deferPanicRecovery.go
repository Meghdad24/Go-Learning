package main

import "fmt"

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

func main() {

	a := []string{"a", "b"}
	print(a, 2)
}

func print(a []string, index int) {
	fmt.Println(a[index])
}
