# [最接近的三数之和](https://leetcode-cn.com/problems/3sum-closest/)

### 题目

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

### 解法

 类似[0015求三数之和](https://github.com/cocowh/algorithm/tree/master/0015_threeSum)

```
func threeSumClosest(nums []int, target int) int {
	ret := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l,r := i + 1,length-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if math.Abs(float64(target - sum)) < math.Abs(float64(target - ret)) {
				ret = sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}
```