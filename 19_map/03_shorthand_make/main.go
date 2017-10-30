package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	//var myGreeting = map[string]string{} <- can also be made like this.
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
