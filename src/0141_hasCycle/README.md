# [环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)

### 题目
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

示例 1：


>输入：head = [3,2,0,-4], pos = 1  
输出：true  
解释：链表中有一个环，其尾部连接到第二个节点。

![图1](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png)

示例 2：

>输入：head = [1,2], pos = 0  
输出：true  
解释：链表中有一个环，其尾部连接到第一个节点。  

![图2](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)

示例 3：

>输入：head = [1], pos = -1  
输出：false  
解释：链表中没有环。

![图3](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)

进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

### 解法

一、hash存遍历节点，存过表明有环，遇nil无环

二、快慢指针，相遇有环，遇nil无环
   
```
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
```