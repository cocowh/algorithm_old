# [分隔链表](https://leetcode-cn.com/problems/partition-list/)

### 题目

给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

>输入: head = 1->4->3->2->5->2, x = 3  
输出: 1->2->2->4->3->5

### 解法

以分隔值区分大小链表，大小链表连接。

```
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
```
