package main

import "fmt"

func main() {
	s := []int{2, 5, 1, 3, 4, 7}
	r := 3
	fmt.Println(shuffle(s, r))
}
func shuffle(nums []int, n int) []int {
	ansSlice := make([]int, 2*n)

	//for i := 0, j := 1; (i <= 2*n -2) && (j <= 2*n - 1); i+2, j+2 {
	//
	//}
	var m, k, j = 0, n, 1
	for i := 0; i <= 2*n-2; i += 2 {

		ansSlice[i] = nums[m]
		ansSlice[j] = nums[k]
		m++
		k++
		j += 2
	}
	return ansSlice
}
