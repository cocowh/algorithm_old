package main

import "fmt"

type ListNode struct {
 	Val int
    Next *ListNode
}

func main()  {
	headA := &ListNode{
		Val:  4,
		Next: nil,
	}
	node1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node8 := &ListNode{
		Val:  8,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	node4.Next = node5
	node8.Next = node4
	node1.Next = node8
	headA.Next = node1
	headB := &ListNode{
		Val:  5,
		Next: nil,
	}
	node0 := &ListNode{
		Val:  0,
		Next: nil,
	}
	node1_2 :=&ListNode{
		Val:  1,
		Next: nil,
	}
	node1_2.Next = node8
	node0.Next = node1_2
	headB.Next = node0
	node := getIntersectionNode(headA,headB)
	fmt.Println(node)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA,pB,hasLinkA,hasLinkB := headA,headB,false,false
	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		}
		pA,pB = pA.Next,pB.Next
		if pA == nil && !hasLinkB {
			pA = headB
			hasLinkB = true
		}
		if pB == nil && !hasLinkA {
			pB = headA
			hasLinkA = true
		}
	}
	return nil
}