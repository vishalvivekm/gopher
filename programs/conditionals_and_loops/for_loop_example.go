package main

import "fmt"

func main() {

	/* infinite loop:
	for {
		fmt.Println("Hi there")
	}
	*/

	i := 3
	for i > 10 {
		fmt.Println(i * 2)
		i += 1
	}
}

// Guess output?
