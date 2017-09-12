package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"

	greeting = append(greeting, "Hej")
	greeting = append(greeting, "Hej")
	greeting = append(greeting, "Hej")
	greeting = append(greeting, "Hej")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
