package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Rickard", 25}
	fmt.Println(&c1.name) // memory address

	changeMe(&c1)

	fmt.Println(c1)       // {Rickard 25}
	fmt.Println(&c1.name) // memory address
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Rickard  25}
	fmt.Println(&z.name) // memory address
	z.name = "Ricky"
	fmt.Println(z)       // &{Ricky 25}
	fmt.Println(&z.name) // memory address

}