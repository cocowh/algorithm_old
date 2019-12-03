# [删除链表的倒数第N个节点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

### 题目

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

### 解法

遍历两次解法：

  * 遍历一便求长度length，再遍历一遍找删除点的父节点，即第length - n - 1个节点，将删除节点的父节点指向删除节点的子节点。
  * 边界：单个节点length - n - 1越界

遍历一次解法：

 * 两个指针遍历，中间间隔n个节点。
 * 第一个指针遍历n个节点后开始一起遍历第二个节点，第一个节点遍历到末尾，第二个节点的位置即为倒数第n个节点的父节点
 * 将删除节点的父节点指向删除节点的子节点

解法一：

```
func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length - 3; i++{
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < length - 2; j++ {
			if j > i +1 && nums[j] == nums[j-1] {
				continue
			}
			l,r := j + 1,length - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target || (l != j + 1 &&  nums[l] == nums[l - 1]){
					l++
				} else if sum > target  || (r != length - 1 && nums[r] == nums[r + 1]){
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return ret
}
```
解法二：

```
func removeNthFromEndTwo(head *ListNode, n int) *ListNode {
	head1,head2,res := head,head,ListNode{}
	res.Next = head
	for head1 != nil {
		head1 = head1.Next
		if n < 0 {
			head2 = head2.Next
		}
		n--
	}
	if n == 0 {	//length = n删除首位
		return head.Next
	}
	head2.Next = head2.Next.Next
	return  res.Next
}
```