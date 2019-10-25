# [旋转数组](https://leetcode-cn.com/problems/rotate-array/)

### 题目

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

>输入: [1,2,3,4,5,6,7] 和 k = 3  
输出: [5,6,7,1,2,3,4]  
解释:  
向右旋转 1 步: [7,1,2,3,4,5,6]  
向右旋转 2 步: [6,7,1,2,3,4,5]  
向右旋转 3 步: [5,6,7,1,2,3,4]  

示例 2:

>输入: [-1,-100,3,99] 和 k = 2  
输出: [3,99,-1,-100]  
解释:   
向右旋转 1 步: [99,-1,-100,3]  
向右旋转 2 步: [3,99,-1,-100]  

说明:

* 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
* 要求使用空间复杂度为 O(1) 的 原地 算法。

### 解法

[参考](https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode/)

```
func rotate(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		cur, pre, cycle := start, nums[start], false
		for start != cur || !cycle {
			cycle = true
			next := (cur + k) % len(nums)
			temp := nums[next]
			nums[next] = pre
			pre = temp
			cur = next
			count++
		}
	}
}

func rotateByReverse(nums []int, k int) {
	k = k % len(nums)
	reverseArr(nums, 0, len(nums)-1)
	reverseArr(nums, 0, k-1)
	reverseArr(nums, k, len(nums)-1)
}

func reverseArr(nums []int, start, end int) {
	for start < end {
		nums[start],nums[end] = nums[end],nums[start]
		start++
		end--
	}
}
```
