package main

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
}
