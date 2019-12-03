# [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

### 题目

将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

### 解法

递归：

```
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return  l1
	} else {
		l2.Next = mergeTwoLists(l1,l2.Next)
		return l2
	}
}
```

迭代插入：

```
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, node *ListNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		head,node = l1,l1
		l1 = l1.Next
	} else {
		head,node = l2,l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 == nil {
		node.Next = l2
	} else {
		node.Next = l1
	}
	return head
}
```