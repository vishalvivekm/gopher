// https://leetcode.com/problems/intersection-of-two-linked-lists
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    mp := make(map[*ListNode]int) 

    temp := headA
    for temp != nil {
        mp[temp] = 1
        temp = temp.Next
    }
     temp1 := headB
    for temp1 != nil {
        _, ok := mp[temp1]
        if ok {
            return temp1
        }
        temp1 = temp1.Next
    }
    
    return nil
}
