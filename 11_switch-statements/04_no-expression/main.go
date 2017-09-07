package main

import "fmt"

func main() {

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Length of 2")
	case myFriendsName == "Mars":
		fmt.Println("Mar")
	default:
		fmt.Println("No match, default.")
	}
}
