 // Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

func containsDuplicate(nums []int) bool {
    Map := make(map[int]int)
    for _, ele := range nums {
        _, found := Map[ele]
        if found {
            return true
        }
       Map[ele] = 1
    }
    return false
}
