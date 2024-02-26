package main

import (
  "fmt"
  "math"
)


func main() {
 arr := []int{1,2,3,4,5}
 k := 2 
 ans := splitArray(arr, k)
 fmt.Println(ans)
}

func splitArray(nums []int, m int) int {
 start := 0.0
 end := 0.0

 for i := 0; i < len(nums); i++ {
 start = math.Max(float64(start), float64(nums[i]))
   end += float64(nums[i])
 }
 
 for start < end {
 mid := start + ( end - start ) / 2

 sum := 0.0
 pieces := 1.0

 for _,ele := range nums {
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

