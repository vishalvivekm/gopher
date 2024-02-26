package main

import (
	"fmt"
)

func keepLastOccurrence(input []int) []int {
	lastOccurrence := make(map[int]int)
	var result []int

	for i := len(input) - 1; i >= 0; i-- {
		if _, found := lastOccurrence[input[i]]; !found {
			result = append(result, input[i])
			lastOccurrence[input[i]] = i
		}
	}

	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {
	// Enter your code here. Read input from STDIN. Print output to STDOUT
	var x, y int
	fmt.Scan(&x)

	var a []int
	for i := 0; i < x; i++ {
		fmt.Scan(&y)
		a = append(a, y)
	}

	elem := keepLastOccurrence(a)
	fmt.Println(len(elem))

	for i := 0; i < len(elem); i++ {
		fmt.Print(elem[i], " ")
	}
	fmt.Println()
}
