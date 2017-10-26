package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":	"Good morning!",
		"one":	"Bonjour!",
		"two":	"Buenos dias!",
		"three":	"Bongiorno!",
	}

	fmt.Println(myGreeting)

	delete(myGreeting, "two")

	if val, exists := myGreeting["two"]; exists {
		fmt.Println("That value exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}
