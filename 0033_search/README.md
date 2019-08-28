# [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

### 题目
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

### 解法

参考: [解法](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-by-leetcode/)
    
```
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0  {
		return -1
	}
	if length == 1{
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	pointIndex := 0
	if nums[0] > nums[length - 1] {
		for left,right := 0,length - 1;left <= right; {
			mid := (left + right) / 2
			if nums[mid] > nums[mid + 1] {
				pointIndex = mid + 1
				break
			} else if nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if nums[pointIndex] == target {
		return pointIndex
	} else if pointIndex == 0 {
		return binarySearch(nums,target,0,length - 1)
	} else if nums[0] > target {
		return binarySearch(nums, target,pointIndex + 1, length -1)
	} else {
		return binarySearch(nums, target, 0, pointIndex - 1)
	}
}

func binarySearch(nums []int, target,left,right int) int {
	for  left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
```