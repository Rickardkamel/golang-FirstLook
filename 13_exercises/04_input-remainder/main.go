package main

import "fmt"

var largeNumber int
var smallNumber int

func main() {
	fmt.Print("Please write a large number: ")
	fmt.Scan(&largeNumber)

	fmt.Print("Please write a small number: ")
	fmt.Scan(&smallNumber)

	fmt.Println(largeNumber, "/",
	smallNumber, "=", largeNumber / smallNumber)
}
