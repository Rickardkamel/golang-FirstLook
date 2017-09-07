package main

import "fmt"

var userInput string

func main() {
	fmt.Print("Please write your name: ")
	fmt.Scan(&userInput)
	fmt.Println("Hello", userInput)
}
