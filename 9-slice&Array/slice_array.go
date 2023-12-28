package main

import "fmt"

func main() {
	var arrayInts = [7]int{1, 25, 12354, 654, 32}
	fmt.Println(arrayInts, len(arrayInts), cap(arrayInts))
}
