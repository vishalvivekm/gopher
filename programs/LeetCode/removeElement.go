// https://leetcode.com/problems/remove-element/
// Beats 100.00%of users with Go

func removeElement(nums []int, val int) int {
	k := 0
	for _, value := range nums {
		if value != val {
			k++
		}
	}

	// traverse array, if ele == val, swap with next element != val, continue

	for i := 0; i < len(nums) && i < k; i++ {
		if nums[i] == val {
			nextIndex := findNextElementIndexNotEqualtoVal(nums, val, i)
			swap(nums, i, nextIndex)
		}
	}
	return k
}
func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func findNextElementIndexNotEqualtoVal(nums []int, val, start int) int {
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			return i
		}
	}
	return -1
}
