package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Daniel":
		fmt.Println("Whats up Daniel")
	case "Marcus":
		fmt.Println("Whats up Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Whats up Medhi")
	default:
		fmt.Println("No friends")
	}
}


