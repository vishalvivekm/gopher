package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arr[1:8]       // array[a:b], including a but excluding b, length of the slice = b-a
	fmt.Println(slice1)      // [2 3 4 5 6 7 8]
	fmt.Println(len(slice1)) // 7
	fmt.Println(cap(slice1))

	subSlice := slice1[0:5]
	fmt.Println(subSlice)

	//arr1 := []int{-1, -2}
	//for _, value := range arr1 {
	//	fmt.Println(value)
	//}
}
