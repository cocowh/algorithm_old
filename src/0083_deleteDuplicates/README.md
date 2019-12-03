# [删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)

### 题目

给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

>输入: 1->1->2  
输出: 1->2

示例 2:

>输入: 1->1->2->3->3  
输出: 1->2->3

### 解法

前后指针

```
type ListNode struct {
     Val int
     Next *ListNode
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
```