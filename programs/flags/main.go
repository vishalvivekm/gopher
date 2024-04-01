package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Check that our program takes exactly three arguments:
	if len(os.Args) != 4 {
		log.Fatal("Error! Expected 3 arguments only!")
	}

	fmt.Printf("Contents of the os.Args slice = %v\n", os.Args)
	fmt.Printf("Name of our program --> os.Args[0] = %[1]s | type: %[1]T\n", os.Args[0])
	fmt.Printf("First cmd-line argument --> os.Args[1] = %[1]s | type: %[1]T\n", os.Args[1])
	fmt.Printf("Second cmd-line argument --> os.Args[2] = %[1]s | type: %[1]T\n", os.Args[2])
	fmt.Printf("Third cmd-line argument --> os.Args[3] = %[1]s | type: %[1]T\n", os.Args[3])
	// values inside os.Args are all string
}
