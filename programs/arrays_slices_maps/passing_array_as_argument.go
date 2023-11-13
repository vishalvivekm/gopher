package main

import "fmt"

// // The size of array is part of  its type, The types [10]int and [20]int are distinct
func calcSquare(numbers []int) ([]int, bool) {
	// func calcSquare(numbers [3]int) ([]int, bool) {
	// [3]int and []int are distinct
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
	}
	return squares, true
}

func main() {
	arr := [3]int{10, 20, 15}
	fmt.Println(arr)
	fmt.Println(calcSquare(arr[:]))
	// fmt.Println(calSquare(arr))
}

// output :

// things to notice:
// arr[:] // when passing array as argument
// multiple returns from the function

// output:
// [10 20 15]
// [100 400 225] true

/*
Explaination
In Go, arr[:] is a slice expression that creates a slice referring to the entire array arr.
Here in the code above, arr is an array with elements {10, 20, 15} of type [3]int. When we use arr[:], it creates a slice that includes all the elements of the array.
In Golang, a slice is a dynamically-sized, flexible view into the elements of an array. It is a more versatile and commonly used data type compared to arrays.

So, when we pass arr[:] to the calcSquare function, we are passing a slice of the array, and the function signature is defined to accept a slice (numbers []int). This allows the function to operate on slices of different lengths.

In contrast, if we had defined the function to accept an array (func calcSquare(numbers [3]int)), we would need to pass an array of exactly length 3, and calcSquare(arr[:]) would not work because the lengths would not match.

Using slices in function arguments is often preferred in Go because it allows for more flexibility in handling arrays of different lengths.
*/
