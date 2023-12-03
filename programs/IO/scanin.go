package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		firstName  string
		secondName string
		lastName   string
	)
	fmt.Println("Enter your first, second and last name")
	fmt.Scanln(&firstName, &secondName, &lastName)
	// fmt.Scanf("%s%s%s", &firstName, &secondName, &lastName)
	fullName := []string{firstName, secondName, lastName}

	fmt.Println("You are: ", strings.Join(fullName, " "))

}
