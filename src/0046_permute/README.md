# [全排列](https://leetcode-cn.com/problems/permutations/)

### 题目

给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

>输入: [1,2,3]  
输出:  

```
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

### 解法

参考[回溯](https://leetcode-cn.com/problems/permutations/solution/quan-pai-lie-by-leetcode/)

```
func backtrack(length, first int, ret *[][]int, nums []int) {
	if first == length {
		tmp := make([]int,length)
		copy(tmp,nums)
		*ret = append(*ret, tmp)
	}
	for i := first; i < length; i++ {
		nums[first],nums[i] = nums[i],nums[first]
		backtrack(length,first+1,ret,nums)
		nums[first],nums[i] = nums[i],nums[first]
	}
}

func permute(nums []int) [][]int {
	ret := [][]int{}
	numsTemp := nums
	backtrack(len(numsTemp), 0, &ret, numsTemp)
	return ret
}
```
