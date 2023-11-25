package main

import "fmt"

type Student struct {
	name string
}

func printChangedName(student *Student) {
	name1 := student.name
	fmt.Println(name1)
}

func printString(s *string) {
	*s = "world"
	fmt.Println(*s, s)
}
func main() {
	s := "hello"
	printString(&s)

	student := Student{"Vivek"}
	printChangedName(&student)
}

// question
/*
why
*/
