# [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

### 题目
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

### 解法

简单难度。

* 记连续和sum，最大连续和maxSum，每次求得连续和与maxSum比较取大值赋给maxSum。
* 若sum > 0，表示连续和目前还不会对后续和产生负增益，还可以继续保留为后续做"贡献"
* 若sum < 0，表示当前连续和会减少后续和的值，应当舍弃，将当前连续和设为当前值。
   
```
func maxSubArray(nums []int) int {
	ret,sum,lenN := nums[0],0,len(nums)
	for i := 0; i < lenN; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > ret {
			ret = sum
		}
	}
	return ret
}
```