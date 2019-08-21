package _025_reverseKGroup

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	ret := ListNode{}
	ret.Next = head
	return ret.Next
}
