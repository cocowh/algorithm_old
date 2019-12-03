package _086_partition

type ListNode struct {
     Val int
     Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	bigList,smallList := &ListNode{},&ListNode{}
	smallTail,bigTail := smallList,bigList
	for head != nil {
		if head.Val < x {
			smallTail.Next = head
			smallTail = smallTail.Next
		} else {
			bigTail.Next = head
			bigTail = bigTail.Next
		}
		head = head.Next
	}
	smallTail.Next = bigList.Next
	bigTail.Next = nil
	return smallList.Next
}