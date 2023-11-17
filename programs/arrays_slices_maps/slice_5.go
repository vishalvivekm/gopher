package main

import (
	"fmt"
	"reflect"
)

func main() {

	// copying from a slice
	//func copy(dst, src []Type) int
	//num := copy(dest_slice, src_slice) // num - how many elements have been copied

	slice := []int{10, 20, 90, 70, 60}
	fmt.Println(reflect.TypeOf(slice)) // []int type

	slice1 := make([]int, 10)
	num1 := copy(slice1, slice)
	fmt.Println(slice1) // [10, 20, 90, 70, 60]
	fmt.Println(num1)

	/*The copy function only copies the minimum of the length of the source and destination slices.*/

	slice2 := make([]int, 3) // declare a slice of int datatype and length 3
	num2 := copy(slice2, slice)
	fmt.Println(slice2) // [10, 20, 90]
	fmt.Println(num2)

}

// to be learnt later
