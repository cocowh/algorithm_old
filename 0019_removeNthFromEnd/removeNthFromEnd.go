package _019_removeNthFromEnd

type ListNode struct {
      Val int
      Next *ListNode
 }
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := ListNode{}
	res.Next = head
	length := 0
	temp := head
	for temp != nil  {
		length++
		temp = temp.Next
	}
	length = length - n
	if  length == 0 {
		return  head.Next
	}
	temp = head
	length--
	for length > 0  {
		temp = temp.Next
		length--
	}
	temp.Next = temp.Next.Next
	return res.Next
}

func removeNthFromEndTwo(head *ListNode, n int) *ListNode {
	head1,head2,res := head,head,ListNode{}
	res.Next = head
	for head1 != nil {
		head1 = head1.Next
		if n < 0 {
			head2 = head2.Next
		}
		n--
	}
	if n == 0 {	//length = n删除首位
		return head.Next
	}
	head2.Next = head2.Next.Next
	return  res.Next
}