package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{10, 20, 90, 70, 60}
	slice := arr[:3]
	fmt.Println(cap(slice))
	newSlice := append(slice, 100, 200)
	fmt.Println(cap(newSlice))

	//newNewSlice := append(slice, 34, 45)
	//fmt.Println(newNewSlice)
	//fmt.Println(cap(newNewSlice)) // 5
	//fmt.Println(cap(slice)) // 5

	newNewSlice := append(slice, 34, 45, 56)
	fmt.Println(newNewSlice)
	fmt.Println(cap(newNewSlice)) // 10
	fmt.Println(cap(slice))       // 5

	// appending to anther slice

	firstArray := [...]int{23, 3, 4, 56, 36, 68}
	fmt.Println(reflect.TypeOf(firstArray)) // [6]int
	firstSlice := firstArray[:2]

	secondArray := [...]int{34, 23, 78, 30}
	secondSlice := secondArray[:3]

	firstSlice = append(firstSlice, secondSlice...) // notice the three dots with the name of second slice
	fmt.Println(firstSlice)

}

// func append(s []T, vs ...T) []T //takes as arguments a slice ( of a certain datatype) and the elements ( of the same datatype)
//  returns a slice containing all the elements of original slice + the passed values
//slice = append(slice, element-1, element-2)
//slice = append(slice, 10, 20, 30)

// if number of elements passed as argument to append into a slice exceeds the capacity, a array with a larger capacity (usually double the current capacity).
//is allocated. and the returned slice has two times the capacity.

// appending to other slice :
// slice = append( slice, anotherSlice...)
/*
fistArray := [...]int{23, 3, 4, 56, 36,68}
*/
