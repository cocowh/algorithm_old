# [三数之和](https://leetcode-cn.com/problems/3sum/)

### 题目

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

### 解法

 [link](https://leetcode-cn.com/problems/3sum/solution/three-sum-ti-jie-by-wonderful611/)

```
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	length := len(nums)
	if length == 0 || nums[0] > 0 || nums[length - 1] < 0 {
		return ret
	}
	for i := 0; i < length; i++ {
		if nums[i] > 0{
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l,r := i + 1,length-1
		for l < r {
			if nums[l] + nums[r] == -nums[i]{
				if l == i + 1 || r == length - 1 {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
					l++
					r--
				} else if nums[l] == nums[l - 1] {
					l++
				} else if nums[r] == nums[r + 1] {
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
					l++
					r--
				}
			} else if nums[l] + nums[r] < -nums[i] {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}
```