package main

import (
	"fmt"
	"io"
	"reflect"
)

type MyStruct struct {
	Field int
}

// implementing io.Reader by MyStruct
func (ms *MyStruct) Read(p []byte) (n int, err error) {
	fmt.Println("implemented io.Reader")
	return 0, nil
}

//func (ms *MyStruct) Write(p []byte) (n int, err error) {
//	fmt.Println("implemented io.Reader")
//	return 0, nil
//}

func main() {
	var ptr *MyStruct
	typ := reflect.TypeOf(ptr)

	//	result := reflect.TypeOf(strings.NewReader("Vivek").Implements(reflect.TypeOf((*io.Reader)(nil)).Elem())
	result := reflect.TypeOf(ptr).Implements(reflect.TypeOf((*io.Reader)(nil)).Elem())
	//	result := reflect.TypeOf(ptr).Implements(reflect.TypeOf((*io.Writer)(nil)).Elem())

	fmt.Println(result)

	// Using .Elem() to get the underlying type of the pointer
	elemType := typ.Elem()

	fmt.Println("Type of ptr:", typ)          //  *main.MyStruct
	fmt.Println("Underlying type:", elemType) //
}
