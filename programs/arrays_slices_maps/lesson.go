package main

import "fmt"

// go makes it possible to return multiple values from a  function
//func operation(a int, b int) (int, int) { // function signature
//	sum := a + b
//	diff := a - b
//	return sum, diff // terminating statement
//}

func operation(a int, b int) (sum int, diff int) { // return values already declared here in func signature, so #z
	sum = a + b
	diff = a - b
	// sum = a + b : order doesn't make a difference also note how we aren't using shorthand op here
	return // #z it eliminates the need of mentioning the variable names from the return statement here
}
func main() {
	sum, difference := operation(20, 10) // function returning two values, which are being stored in sum, difference
	fmt.Println(sum, difference)
}

//>>> go run main.go
//30 10
