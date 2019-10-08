package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	list := &ListNode{}
	node := list
	for i := 1; i <=5; i++ {
		tempNode := &ListNode{
			Val:  i,
			Next: nil,
		}
		node.Next = tempNode
		node = node.Next
	}
	newlist := rotateRight(list.Next, 10)
	for newlist != nil {
		fmt.Println(newlist.Val)
		newlist = newlist.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length,tailNode := 0,head
	for tailNode.Next != nil {
		length++
		tailNode = tailNode.Next
	}
	length++
	tailNode.Next = head
	for i := 0; i < (length - k % length) % length; i++ {
		tailNode = tailNode.Next
	}
	head = tailNode.Next
	tailNode.Next = nil
	return head
}