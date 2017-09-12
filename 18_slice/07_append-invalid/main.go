package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour"
	greeting[2] = "buenos dias!"
	greeting[3] = "Hej!"

	fmt.Println(greeting[2])
	// greetings[3] cant be appended like this.
}
