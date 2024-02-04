// problem statement

/*Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.

Return the minimized largest sum of the split.

A subarray is a contiguous part of the array.

Given an integer array nums add an integer k, split nums into k non-empty subarrays such that the larges sum of any subarray is minimized.
Return the minimized largest sum of the split.

Test case:

Input: nums = [7,2,5,10,8], k = 2
Output: 18 : correct answer
Explanation: There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8], where the largest sum among the two subarrays is only 18. */

// solution implementation:

package main

import (
 "fmt"
 "math"
)

func main() {
  arr := []int{7,2,5,10,8}
  k := 2
  fmt.Println(splitArray(arr, k))

}


func splitArray(nums []int, m int) int {
 start := 0.0
 end := 0.0

 for i := 0; i < len(nums); i++ {
  start = math.Max(float64(start), float64(nums[i]))
  end += float64(nums[i])}

  for start < end {

   mid := start + ( end - start ) / 2

   sum := 0.0
   pieces := 1.0 // writing these af floats is necessary,as math.Max() in golang takes only float arguemnts, 
   // and I don't know any alternative method right now :(

   for _, ele := range nums {
    
	if sum + float64(ele) > mid {
	sum = float64(ele)
	pieces++
	} else {
      sum += float64(ele)
	}
	}
    if pieces > float64(m) {
      start = mid + 1
	} else {
     end = mid 
	}


  }

return int(end)

}
