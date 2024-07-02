package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	b, err := reader.ReadBytes('\n') // Input into `b`: Hello World!\n
	if err != nil {
		log.Fatal(err) // Exit if we have an unexpected error
	}
	fmt.Println(string(b)) // Output: Hello World!\n, returned string contains the delimiter

	s, err := reader.ReadString('s') // Input into `s`: Vivek Vishal\n
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s) // Output: Vivek Vis
}
