package main

import "fmt"

func main() {
	x := 0
	// anonymous function (func expression - assigning a function to a variable)
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope.
*/