/*Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
Return the array in the form [x1,y1,x2,y2,...,xn,yn].*/

func shuffle(nums []int, n int) []int {
    ansSlice := make([]int, 2*n)
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
