package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "credit"
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}