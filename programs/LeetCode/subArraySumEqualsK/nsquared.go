// https://leetcode.com/problems/subarray-sum-equals-k/

func subarraySum(nums []int, k int) int {
    if(len(nums) == 1) {
        if nums[0] != k {
            return 0
        }
    }
    count := 0
    // create prefix sum array
    parr := createPrefixSumArr(nums)

    for i := 0; i < len(nums) -1 ; i++ {
       if parr[i] == k {
                count = count + 1
       }
       for j := i+1; j < len(nums); j++ {
           
            if parr[j] - parr[i] == k {
             
               count = count + 1
            }
       }

    }
    if parr[len(nums) - 1] == k { // account for parr[len(sum) -1 ] == k
        count++
    }
    return count
}
func createPrefixSumArr(arr []int) []int {
    ans := make([]int, len(arr))
    ans[0] = arr[0]
    for i := 1; i < len(arr); i++ {
       ans[i] = ans[i-1] + arr[i]
        
    }
    return ans
}
