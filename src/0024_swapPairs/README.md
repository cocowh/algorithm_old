# [括号生成](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

### 题目

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

### 解法

递归：
 * 设交换的两个节点为head和next，head连接后面交换完成的子链表，next连接head，完成交换
 * 当无节点或只剩一个节点（head == nil || head.Next == nil）,无法进行交换，终止
     

```
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
```