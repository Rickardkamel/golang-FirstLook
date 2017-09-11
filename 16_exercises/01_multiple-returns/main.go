package main

import "fmt"

func main() {
	h, even := divide(5)
	fmt.Println(h, even)
}

func divide(i int) (int, bool) {
	return i / 2, i%2 == 0
}
