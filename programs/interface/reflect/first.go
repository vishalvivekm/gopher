package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func main() {
	var c *bytes.Buffer
	//d := io.Writer(&bytes.Buffer{})
	//e := io.Reader(&bytes.Buffer{})
	//f := io.Reader(strings.NewReader("hello"))

	//fmt.Printf("%[1]v, %[1]T", d)
	// fmt.Printf("%v, %T\n", d, d)

	// Check if c implements io.Writer
	//ans := reflect.TypeOf(c).Implements(reflect.TypeOf((*io.Writer)(nil)).Elem())
	// fmt.Println(ans)
	if reflect.TypeOf(c).Implements(reflect.TypeOf((*io.Writer)(nil)).Elem()) {
		fmt.Println("c implements io.Writer")
	} else {
		fmt.Println("c does not implement io.Writer")
	}
}
