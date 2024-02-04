

// Contains Duplicate II : Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

import (
    "fmt"
    "math"
)
func containsNearbyDuplicate(nums []int, k int) bool {
        Map := make(map[int]int)
    for index, ele := range nums {
        _, found := Map[ele]
        if found {
            if math.Abs(float64(Map[ele] - index)) <= float64(k) {
                return true
            }
        }
       Map[ele] = index
    }
    return false
}
