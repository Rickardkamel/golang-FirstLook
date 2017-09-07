package main

import "fmt"

/*
switch on types
-- normally switch on value of variables
-- go allows you to switch on type of variable
*/

type Contact struct {
	greeting string
	name     string
}

//Interfaces - later
func SwitchOnType(x interface{}) {
	switch x.(type) { //this is an assert/asserting "x is of this type"

	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Kml")
	var t = Contact{"Good to see you,", " Tommy"}
	SwitchOnType(t)
}
