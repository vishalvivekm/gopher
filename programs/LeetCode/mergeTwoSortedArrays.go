func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, k := 0,0,0
    ans := make([]int, m+n)
    for i < m && j < n {
        if nums1[i] < nums2[j]{
            ans[k] = nums1[i]
            i++ 
            k++
        } else {
            ans[k] = nums2[j]
            j++
            k++
        }
      
    }
    for i < m {
        ans[k] = nums1[i]
        i++
        k++
    }
    for j < n {
        ans[k] = nums2[j]
        j++
        k++
    }
    fmt.Println(ans)
    _ = copy(nums1, ans) 
}
