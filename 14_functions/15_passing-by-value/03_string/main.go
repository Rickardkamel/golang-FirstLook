package main

import "fmt"

func main() {
	name := "Rickard"
	fmt.Println(name) // Rickard

	changeMe(name)

	fmt.Println(name) // Rickard
}

func changeMe(z string) {
	fmt.Println(z) // Rickard
	z = "Ricky"
	fmt.Println(z) // Ricky
}
