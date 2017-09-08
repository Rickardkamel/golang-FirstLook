package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...) // variadic argument
	fmt.Println(n)
}


func average(sf ...float64) float64 { // variadic params
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf { // for (index), (value) := range (variable)
		total += v
	}

	return total / float64((len(sf)))
}