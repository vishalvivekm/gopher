/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// fast and slow pointer approach https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil;
    }
    if head.Next.Next == nil {
        head.Next = nil
        return head
    }
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    value := slow.Next.Val
    slow.Val = value
    slow.Next = slow.Next.Next
    
    return head
}
