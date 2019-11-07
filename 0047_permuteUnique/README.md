# [全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

### 题目
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

>输入: [1,1,2]  
输出:  

```
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```
### 解法

不会做，[参考](https://leetcode-cn.com/problems/permutations-ii/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liwe-2/)。

[参考](https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0047.permutations-ii/permutations-ii.go)。

```
func backtrack(res *[][]int, nums, stack []int, token []bool, length, start int) {
	if start == length {
		temp := make([]int, len(stack))
		copy(temp, stack)
		*res = append(*res, temp)
		return
	}
	used := make(map[int]bool, length-start)
	for i := 0; i < length; i++ {
		if !token[i] && !used[nums[i]] {
			used[nums[i]] = true
			token[i] = true
			stack[start] = nums[i]
			backtrack(res, nums, stack, token, length, start+1)
			token[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	numsTemp := nums
	sort.Ints(numsTemp)
	token := make([]bool, len(numsTemp))
	stack := make([]int,len(numsTemp))
	backtrack(&res, numsTemp, stack, token, len(numsTemp), 0)
	return res
}
```
