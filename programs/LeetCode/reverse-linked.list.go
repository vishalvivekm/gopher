// https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var prev *ListNode
    temp := head
    
    for ;temp.Next != nil; {
        front := temp.Next
        temp.Next = prev
        prev = temp
        temp = front
    }
    temp.Next = prev
    head = temp
    return head
}
