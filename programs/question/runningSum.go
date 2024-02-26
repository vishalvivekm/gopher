package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	runningSumm := make([]int, len(nums))
	n := len(nums)
	if n == 0 {
		runningSumm = nil
	} else {
		i := 0
		for i < n {
			runningSumm[i] = sum(nums[0 : i+1]...)
			i++
		}
	}
	return runningSumm
}

func sum(ele ...int) (summ int) {
	summ = 0
	for _, elements := range ele {
		summ += elements
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(runningSum(nums))
}
