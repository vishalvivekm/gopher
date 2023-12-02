package main

import "fmt"

// Declare variable activeUserCount
// your code goes here
var activeUserCount int

func entry() {
	// Hint: you can use the "++" operator to increment a variable by 1
	// your code goes here
	activeUserCount++
}

func exit() {
	// Hint: you can use the "--" operator to decrement a variable by 1
	// your code goes here
	activeUserCount--
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}

