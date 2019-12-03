package main

import "fmt"

type ListNode struct {
	Val int
    Next *ListNode
}

func main()  {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	two := ListNode{
		Val:  1,
		Next: nil,
	}
	three := ListNode{
		Val:  2,
		Next: nil,
	}
	four := ListNode{
		Val:  3,
		Next: nil,
	}
	five := ListNode{
		Val:  3,
		Next: nil,
	}
	four.Next = &five
	three.Next = &four
	two.Next = &three
	head.Next = &two

	fmt.Println(deleteDuplicates(head))
	for head != nil  {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	f := head
	e := f.Next
	for f.Next != nil {
		if e.Val == f.Val {
			f.Next = e.Next
		} else {
			f = f.Next
		}
		e = e.Next
	}
	return head
}
