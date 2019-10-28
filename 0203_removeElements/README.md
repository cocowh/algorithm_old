# [移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements/)

### 题目

删除链表中等于给定值 val 的所有节点。

示例:

>输入: 1->2->6->3->4->5->6, val = 6  
输出: 1->2->3->4->5

### 解法

* 确定head.val != val，遍历head至新head
* 遍历head，确定head.Next.Val != val

```
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	if head == nil {
		return head
	}
	point := head
	for point.Next != nil {
		if point.Next.Val == val {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
		if point == nil {
			break
		}
	}
	return head
}
```
