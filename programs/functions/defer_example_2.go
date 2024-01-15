/*
StackOverFlow link: https://stackoverflow.com/questions/77521665/when-exactly-is-the-deferred-statement-is-executed
*/

package main

import "fmt"

func example() (result int) {
	defer func() {
		result = result * 2 // resultt := result * 2
		//fmt.Println(resultt)
	}()
	fmt.Println("Executing main logic")
	return 42
}

func main() {
	fmt.Println(example())
}

// output?
// Executing main logic
// 84

/*

golang.org/ref/spec#Defer_statements "if the surrounding function returns through an explicit return statement,
deferred functions are executed after any result parameters are set by that return statement but before the function returns to its caller."
... "if the deferred function is a function literal and the surrounding function has named result parameters that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned."
*/

//icza:  Deferred functions are executed before returning a value, so deferred functions can modify the returned values (if they are named, because you have to refer to them)
