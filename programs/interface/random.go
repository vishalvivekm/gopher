package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func main() {
	var I io.Writer
	var b bytes.Buffer  // doesn't implement io.Writer
	var c *bytes.Buffer // implements io.Writer

	//I = b // wrong
	//I = c // correct

	fmt.Println(reflect.TypeOf(c))
	//fmt.Fprintln(I, "Vivek Vishal")

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(I)

	d := interface{}(c)
	val, ok := d.(io.Writer) // type assertion

	if ok {
		fmt.Println("d implements io.Writer: ", val)
	} else {
		fmt.Println("d doesn't implement io.Writer")
	}
	e := interface{}(b)
	val1, okk := e.(io.Writer)

	if okk {
		fmt.Println("d implements io.Writer: ", val1)
	} else {
		fmt.Println("e doesn't implement io.Writer")
	}

	i := interface{}(4)
	fmt.Printf("%v is of type %T\n", i, i)

}

// how to make any value a interface
// i := interface{}(yourvalue)
