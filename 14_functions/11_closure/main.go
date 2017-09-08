package main

import "fmt"

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
