package main

import "fmt"

func SumArray(arr [5]int) int {
	sum := 0

	// iterate all elements
	for k := 0; k < len(arr); k++ {
		sum += arr[k]
	}

	return sum
}

func main() {

	// initial an array
	var arr = [5]int{12, 58, 7, 42, 79}
	var sum int

	// Passing an array as argument
	sum = SumArray(arr)
	fmt.Printf("Sum of the array: %d\n", sum)
}

/*
package main

import "fmt"

func SumArray(arr []int) int {
	sum := 0

	// iterate all elements
	for k := 0; k < len(arr); k++ {
		sum += arr[k]
	}

	return sum
}

func main() {

	// initial an array
	var arr = []int{12, 58, 7, 42, 79, 23, 24}
	var sum int

	// Passing an array as argument
	sum = SumArray(arr[:])
	fmt.Printf("Sum of the array: %d \n", sum)
}

*/
