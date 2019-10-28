package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	point := head
	for i := 1; i <= 6; i++ {
		temp := &ListNode{
			Val:  i,
			Next: nil,
		}
		point.Next = temp
		point = point.Next
	}
	ret := removeElements(head.Next, 6)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	if head == nil {
		return head
	}
	point := head
	for point.Next != nil {
		if point.Next.Val == val {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
		if point == nil {
			break
		}
	}
	return head
}
