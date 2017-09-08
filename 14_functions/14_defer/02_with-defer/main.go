package main

import (
	"fmt"
)

func main() {
	defer world() // defers to right before main exits
	hello()
}

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}