package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if even(i) {
			fmt.Println(i)
		}
	}
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}