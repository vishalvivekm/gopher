package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("../temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nHope you had a good day!")
	if err != nil {
		fmt.Println(err)
	}
}

// WriteString is like Write, but writes the contents of string s rather than a slice of bytes
