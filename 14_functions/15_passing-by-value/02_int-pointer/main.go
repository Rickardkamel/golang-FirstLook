package main

import "fmt"

func main() {
	age := 25
	fmt.Println(&age) // memory address

	changeMe(&age) // pass memory address

	fmt.Println(&age)
	fmt.Println(age) // 24
}

func changeMe(z *int) {
	fmt.Println(z)  // memory address
	fmt.Println(*z) // 25
	*z = 24
	fmt.Println(z)  // same memory address
	fmt.Println(*z) // 24
}
