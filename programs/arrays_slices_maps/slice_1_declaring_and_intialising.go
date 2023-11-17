package main

import "fmt"

func main() {
	// declaring and initialising a slice
	// method 1
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arr[1:7] // array[a:b], including a but excluding b, length of the slice = b-a,
	// [:b] - from the begining till b ( excluding b)
	// [a:] - from the a ( including a) till the very end hehe
	fmt.Println(cap(slice1)) // 9
	fmt.Println(slice1)      // [2 3 4 5 6 7 8]
	fmt.Println(len(slice1)) // 7
	fmt.Println(cap(slice1))

	subSlice := slice1[0:5]
	fmt.Println(subSlice)

	// method 2
	// grades := []int{10, 20, 30}
	//arr1 := []int{-1, -2, -3, -4}
	//for _, value := range arr1 {
	//	fmt.Println(value)
	//}

	// remember [4]int and [5]int and []int are different data types in golang [] is part of the type
	// Method 3 : using make([]int, length, capacity)
	slice3 := make([]int, 4, 8) // a
	fmt.Println(slice3)
	for index, _ := range slice3 {
		slice3[index] = index // notive we won't use := here as slice3 has already be defined to have int values
		// fmt.Printf("%d, ", slice3[index])
	}
	fmt.Println(slice3)
}

// for array length = capacity
// for slices, cap is elements from the begining of the slice till the end of the underlying array
// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// slice1 := arr[1:7] // array[a:b], including a but excluding b, length of the slice = b-a,
// [:b] - from the begining till b ( excluding b)
// [a:] - from the a ( including a) till the very end hehe
// fmt.Println(cap(slice1)) // 9
