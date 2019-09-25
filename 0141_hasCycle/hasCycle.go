package _141_hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleByDict(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head !=  nil {
		if _,ok := m[head];ok {
			return true
		} else {
			m[head] = true
		}
		head = head.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	fast,slow := head,head
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		} else {
			return false
		}
	}
	return  false
}

