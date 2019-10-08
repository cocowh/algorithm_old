# [旋转链表](https://leetcode-cn.com/problems/rotate-list/)
### 题目
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

>输入: 1->2->3->4->5->NULL, k = 2  
输出: 4->5->1->2->3->NULL  
解释:  
向右旋转 1 步: 5->1->2->3->4->NULL  
向右旋转 2 步: 4->5->1->2->3->NULL  

示例 2:

>输入: 0->1->2->NULL, k = 4  
输出: 2->0->1->NULL   
解释:  
向右旋转 1 步: 2->0->1->NULL  
向右旋转 2 步: 1->2->0->NULL  
向右旋转 3 步: 0->1->2->NULL  
向右旋转 4 步: 2->0->1->NULL  


### 解法

* 扫描一遍，确定链表长度length和链表尾tailNode
* tailNode.Next = head链表尾指向表头形成环
* (length - (k % length)) % length确定寻找断点前一个节点需要移动tailNode的次数
* 移动tailNode寻找断点前一节点
* 此时tailNode.Next即为新的头节点，将tailNode.Next指向nil断开环

```
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length,tailNode := 0,head
	for tailNode.Next != nil {
		length++
		tailNode = tailNode.Next
	}
	length++
	tailNode.Next = head
	for i := 0; i < (length - k % length) % length; i++ {
		tailNode = tailNode.Next
	}
	head = tailNode.Next
	tailNode.Next = nil
	return head
}
```