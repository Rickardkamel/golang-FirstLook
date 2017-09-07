package main

import "fmt"

func main() {
	if false {
		fmt.Println("first print")
	} else if false {
		fmt.Println("second print")
	} else if true{
		fmt.Println("third print")
	} else {
		fmt.Println("fourth print")
	}
}
