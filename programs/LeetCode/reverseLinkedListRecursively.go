package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}
/* func (l *ListNode) NewNode(val int, listNode *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: listNode,
	}
} */

//type Head *ListNode
//type Tail *ListNode
 type LinkedList struct {
	head *ListNode // Head
	tail *ListNode // Tail
}

func (l *LinkedList) Insert(val int) {
	node := NewNode(val)
	if l.head == nil {
		l.head = node
		l.tail = l.head  // node
		return
	}

	l.tail.Next = node
	l.tail = node

}
func NewLL() *LinkedList{
return &LinkedList{
head: nil,
tail: nil,
}
}
func (l *LinkedList) print() {
	temp := l.head
	// i := 1
	for temp != nil {
		// fmt.Printf("%d -> %d\n", i, temp.Val)
		fmt.Printf("%d -> ",temp.Val)
		//i++
		temp = temp.Next
	}
	fmt.Println("X")

}
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	return r(head, nil)
}

func r(node *ListNode, prev *ListNode) *ListNode {
	if node.Next == nil {
		node.Next = prev
		prev = node
		return prev
	}
	t := node.Next
	node.Next = prev
	prev = node
	return r(t, prev)
}

func main() {
	s := []int{1, 2, 3, 4}
	var list =  NewLL()

	for _, val := range s {
		list.Insert(val)
	}
	list.print()

}
