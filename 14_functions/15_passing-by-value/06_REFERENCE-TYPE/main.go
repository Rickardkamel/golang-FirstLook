package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Rickard]
}

func changeMe(z []string) {
	z[0] = "Rickard"
	fmt.Println(z) // Rickard
}
