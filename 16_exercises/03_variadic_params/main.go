package main

import "fmt"

func main() {
	largest := max(24, 23, 14, 22, 53, 23, 66, 292)
	fmt.Println(largest)
}

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}
