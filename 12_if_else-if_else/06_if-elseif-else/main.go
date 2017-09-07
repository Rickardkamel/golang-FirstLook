package main

import "fmt"

func main() {
	if false {
		fmt.Println("first print")
	} else if true {
		fmt.Println("second print")
	} else {
		fmt.Println("third print")
	}
}
