# [四数之和](https://leetcode-cn.com/problems/4sum/)

### 题目

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

### 解法
 同三数和，外加一层遍历

 开始自己思路求所有两数和sum存字典map[sum][][]int，map[target - sum]存在则取字典记录遍历合并作为结果，之后再遍历结果排序合并为字符串，存mapT，去重。
 时间复杂度O(n^2)，需要额外存储空间。

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