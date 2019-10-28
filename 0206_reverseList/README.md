# [反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

### 题目

反转一个单链表。

示例:

>输入: 1->2->3->4->5->NULL  
输出: 5->4->3->2->1->NULL  

进阶:  
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

### 解法

```

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

func reverseListSecond(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
```
