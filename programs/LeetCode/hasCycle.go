// https://leetcode.com/problems/linked-list-cycle
func hasCycle(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }   
    return false
}
