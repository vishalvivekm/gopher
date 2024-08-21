// https://leetcode.com/problems/merge-two-sorted-lists/
// O(n) and O(1)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    node := &ListNode{}
    head := node
    temp1 := list1
    temp2 := list2

    for temp1 != nil && temp2 != nil {
        if temp1.Val > temp2.Val {
            node.Next = temp2
            temp2 = temp2.Next
            node = node.Next
        } else {
            node.Next = temp1
            temp1 = temp1.Next
            node = node.Next
        }
    }

    if temp1 == nil {
        node.Next = temp2
    }
    if temp2 == nil {
        node.Next = temp1
    }

    return head.Next
}

// Java : faster than 100% java online submissions :lol
class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        
        ListNode node = new ListNode();
        ListNode head = node;
        // if ( node == null) {
        //     System.out.println("yes both head and node are null at this moment");
        // }
        ListNode temp1 = list1;
        ListNode temp2 = list2;
        
        // if (temp1 == null) {
        //     System.out.println("temp1 is null");
        // }
        while (temp1 != null && temp2 != null) {
            if(temp1.val > temp2.val) {
                node.next = temp2;
                temp2 = temp2.next;
                node = node.next;
            } else {
                node.next = temp1;
                temp1 = temp1.next;
                node = node.next;
            }
        }
       if( temp1 == null) {
           node.next = temp2;
       }
        if(temp2 == null) {
            node.next = temp1;
        }
        return head.next;
    }
}
/*
Test Case: 
[]
[]
*/
