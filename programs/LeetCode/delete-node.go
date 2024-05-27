// https://leetcode.com/problems/delete-node-in-a-linked-list/
// 
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    value := node.Next.Val
    node.Val = value
    node.Next = node.Next.Next
}
