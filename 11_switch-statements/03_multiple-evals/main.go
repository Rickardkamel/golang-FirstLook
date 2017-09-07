package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Hello Tim & Jenny")
	case "Jane", "John":
		fmt.Println("Hello Tim & Jenny")
	}
}
