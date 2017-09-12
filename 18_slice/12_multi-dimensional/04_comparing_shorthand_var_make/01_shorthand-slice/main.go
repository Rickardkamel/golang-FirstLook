package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	//student[0] = "Todd" <- won't work (don't have index pos yet)
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
