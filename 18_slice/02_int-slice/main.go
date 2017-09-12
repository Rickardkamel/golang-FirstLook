package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 8, 11}

	for i, value := range xs {
		fmt.Println(i, " - ", value)
	}
}
