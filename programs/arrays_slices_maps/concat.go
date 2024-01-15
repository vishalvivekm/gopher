package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	//fmt.Printf("%T", arr)
	//ans := append(arr, arr...)

	ans := getConcatenation(arr)
	fmt.Println(ans)
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
