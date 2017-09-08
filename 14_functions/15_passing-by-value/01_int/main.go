package main

import "fmt"

func main() {
	age := 25
	changeMe(age)    // pass value
	fmt.Println(age) // 25
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

/*
When changeMe is called
the value 44 is being passed as an argument
*/
