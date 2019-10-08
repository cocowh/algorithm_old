# [反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)
### 题目
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

>输入: 1->2->3->4->5->NULL, m = 2, n = 4  
输出: 1->4->3->2->5->NULL

### 解法

* 扫描一遍，将开始反转点和终止反转点之间的节点入栈，确定开始反转点和终止反转点。
* 出栈，将节点追加到开始反转点，最后链接终止反转点。
* 注意从开头开始反转时，链表头的确定。

```
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	headNode,tailNode,count,stack := head,head,0,[]*ListNode{}
	for count < n {
		count++
		if count >= m && count <= n {
			stack = append(stack,tailNode)
		}
		if count < m - 1 {
			headNode = headNode.Next
		}
		if count < n + 1 {
			 tailNode = tailNode.Next
		}
	}
	nodeNum := len(stack)
	for i := nodeNum - 1; i >= 0; i-- {
		headNode.Next = stack[i]
		if m == 1 && i == nodeNum - 1 {  //特殊点从开头开始反转，链表头改变
			head = headNode.Next
		}
		headNode = headNode.Next
	}
	headNode.Next = tailNode
	return head
}
```

