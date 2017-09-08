package main

import "fmt"

func main() {

	name := "Rickard"
	fmt.Println(&name) // memory address

	changeMe(&name)

	fmt.Println(&name) // memory address
	fmt.Println(name)  // Ricky
}

func changeMe(z *string) {
	fmt.Println(z)  // memory address
	fmt.Println(*z) // Rickard
	*z = "Ricky"
	fmt.Println(z)  // memory address
	fmt.Println(*z) // Ricky
}
