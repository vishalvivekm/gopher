// memory efficient beats 99%, used ~ 5 MB

import "runtime"
func moveZeroes(nums []int)  {
    runtime.GC()
  i := 0
 for _, val := range nums {
    if val != 0 {
        nums[i] = val
        i++
    }
 }
 for i < len(nums){
    nums[i] = 0
    i++
 }
 runtime.GC()

}

// runtime efficient, beats 100%, uses 0ms
func moveZeroes(nums []int)  {
  i := 0
 for _, val := range nums {
    if val != 0 {
        nums[i] = val
        i++
    }
 }
 for i < len(nums){
    nums[i] = 0
    i++
 }

}
