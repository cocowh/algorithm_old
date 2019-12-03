package _234_isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	nums := make([]int, 0, 32)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	h, t := 0, len(nums)-1
	for h < t {
		if nums[h] != nums[t] {
			return false
		}
		h++
		t--
	}
	return true
}
