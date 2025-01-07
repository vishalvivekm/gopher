func majorityElement(nums []int) int {
    count := 0
   // n := len(nums)
    candidate := 0

    for _, val := range nums{
        if count == 0 {
            candidate = val
        }
        if candidate == val {
            count++
        } else {
            count--
        }
    }
    return candidate
}
