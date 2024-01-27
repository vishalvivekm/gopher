// Given an integer array nums, return the rotated array nums to the right by k steps, where k is non-negative.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotate(arr, 3))
}
func rotate(nums []int, k int) []int {
	n := len(nums)
	nums = append(nums[n-k:n], nums[0:n-k]...)
	// fmt.Println(nums)
	return nums

}
