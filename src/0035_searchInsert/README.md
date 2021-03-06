# [搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/)

### 题目
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

### 解法

二分查找左右收缩，注意nums[mid] == target时
    
```
func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length  == 0 || nums[0] > target{
		return 0
	} else if nums [length - 1]  < target {
		return length
	}
	left,right := 0,length - 1
	for left <= right {
		if nums[left] == nums[right] {
			if nums[left] >= target {
				return left
			} else {
				return left + 1
			}
		}
		if nums[left] ==target {
			return left
		}
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] == target {
			left++
		}
	}
	if nums[left] >= target {
		return left
	} else {
		return left + 1
	}
}
```