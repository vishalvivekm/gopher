package main

import "fmt"

// we can initiate the panic event directly. With the panic() function, we can terminate the execution by ourselves. The panic() function can take any data type as an argument and then print it to the output:
func main() {
	defer fmt.Println("will be printed anyway") // The defer statement works! The panic() func has no effect on the list of deferred calls (if we declare it before the panic).
	panic("Something has gone wrong!")
	// defer fmt.Println("will it be printed ?") // nope
	// fmt.Println("Not printed because of the panic")
}

// Output:
// panic: Something has gone wrong!
/*package main

import "fmt"

func main() {
    defer fmt.Println("Will be printed anyway!")
    panic("Something has gone wrong!")
}

// Output:
// Will be printed anyway!
// panic: Something has gone wrong!*/
