package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Whats up Daniel")
	case "Medhi":
		fmt.Println("Whats up Medhi")

	default:
		fmt.Println("No friends")
	}
}

// No default fallthrough (optional)

