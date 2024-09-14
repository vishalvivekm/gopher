// https://leetcode.com/contest/weekly-contest-414/problems/reach-end-of-array-with-max-score/

func findMaximumScore(nums []int) int64 {
    start := nums[0]
    index := 0
    score := 0
    
    for i, val := range nums {
        if val > start || i == len(nums) - 1{
            score = score + ((i-index) * start)
            index = i
            start = val
        }
    }
    if score == 0 {
        score = len(nums) - 1 
    }
    return int64(score)
}
