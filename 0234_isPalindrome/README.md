# [回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/)

### 题目

请判断一个链表是否为回文链表。

示例 1:

>输入: 1->2  
输出: false

示例 2:

>输入: 1->2->2->1  
输出: true

进阶：  
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

### 解法

容易想到使用数组存遍历值，然后对数组操作校验。不满足进阶条件。

关于进阶解法[参考](https://leetcode-cn.com/problems/palindrome-linked-list/solution/javashi-xian-kuai-man-zhi-zhen-fan-zhuan-qian-ban-/)。

* 先将前一半反转，得到中间节点
* 前后半部分链表比较

改变了链表结构。

```
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	nums := make([]int, 0, 32)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	h, t := 0, len(nums)-1
	for h < t {
		if nums[h] != nums[t] {
			return false
		}
		h++
		t--
	}
	return true
}
```
