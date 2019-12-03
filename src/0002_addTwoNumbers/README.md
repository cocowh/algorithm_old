# [两数相加](https://leetcode-cn.com/problems/add-two-numbers/)
### 题目

给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

### 解法

题意：返回链表节点值为l1，l2每次遍历对应节点值、及上一次相加的进位值之和除10的余数

遍历终止条件：l1和l2是否遍历完（下个节点指向nil）或者进位是否大于0


```
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	retListNode := &ListNode{}  //结果链表
	curNode := retListNode 			//当前节点
	carry := 0					//进位

	for  l1 != nil || l2 != nil || carry > 0  { //未计算完或者有进位
		sum := carry			//相加结果值初始化未进位值
		if l1 != nil {			//链表链链l1未遍历完时
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {			//链表l2未遍历完时
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10		//计算进位
		val := sum % 10			//计算当前值
		curNode.Val = val       	//设置当前节点值
		if l1 != nil || l2 != nil || carry > 0 { //未计算完或者有进位新增下一个节点
			curNode.Next = &ListNode{}
		} else { //否则指向nil
			curNode.Next = nil
		}
		curNode = curNode.Next			//当前节点指向下一个节点
	}
	return retListNode
}
```