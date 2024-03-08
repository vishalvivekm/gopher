package main

import (
	"fmt"
	"strings"
)

/*
type Reader struct {
s string
i int
prevRune
}
*/
func main() {

	r := strings.NewReader("Read from this okay")
	fmt.Printf("%#v\n", r)

	//name := "vivek"
	//var name1 []string
	//
	//n := copy(name1, name[0:])

	name := "vivek"
	var name1 []byte = make([]byte, len(name))
	// here if we declare name1 like var name1 []byte, it'll be a nil slice: you're trying to copy data into a nil byte slice (name1). we need to initialize name1 with a non-nil byte slice ( using make()) before using copy.

	n := copy(name1, name) // as a special case Copy copies bytes from a string to slice of bytes
	//  n := copy(name1, name[:])
	fmt.Println(name1)
	fmt.Println(string(name1))
	fmt.Println(n)

}
