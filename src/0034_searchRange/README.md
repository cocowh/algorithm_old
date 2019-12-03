# [在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

### 题目
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

### 解法

二分左右收缩至确定左右边界

特殊用例：单个元素

二分特殊情景：nums[mid] == target，左右收缩
    
```
func searchRange(nums []int, target int) []int {
	start,end,length := -1,-1,len(nums)
	if length == 0 || nums[0] > target || nums[length - 1] < target{
		return  []int{start, end}
	} else if length == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{start, end}
		}
	}
	left,right := 0,length - 1
	for left <= right && ( end == -1 || start == -1) {
		if (start != -1 && end != -1) {
			break
		}
		if nums[left] == target  && start == -1{
			start = left
		}
		if nums[right] == target && end == -1 {
			end = right
		}
		mid := (left + right) / 2
		if nums[mid] > target && end == -1 {
			right = mid - 1
		}
		if nums[mid] < target && start == -1 {
			left = mid + 1
		}
		if nums[mid] == target {
			if start == -1 {
				left = left + 1
			}
			if end == -1 {
				right = right - 1
			}
		}
	}
	return  []int{start, end}
}
```