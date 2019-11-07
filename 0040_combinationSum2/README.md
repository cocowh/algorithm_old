# [组合总和II](https://leetcode-cn.com/problems/combination-sum-ii/)

### 题目

给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

* 所有数字（包括目标数）都是正整数。
* 解集不能包含重复的组合。 

示例 1:

>输入: candidates = [10,1,2,7,6,1,5], target = 8,  
所求解集为:

```
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

示例 2:

>输入: candidates = [2,5,2,1,2], target = 5,  
所求解集为:

```
[
  [1,2,2],
  [5]
]
```

### 解法

参考[回溯](https://leetcode-cn.com/problems/combination-sum/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-2/)

```
func findCombinationSum(ret *[][]int, candidates, stack []int, target, start int) {
	if target == 0 {
		temp := make([]int, len(stack))
		copy(temp, stack)
		*ret = append(*ret, temp)
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		stack = append(stack, candidates[i])
		findCombinationSum(ret, candidates, stack, target-candidates[i], i+1)
		stack = stack[:len(stack)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	findCombinationSum(&res, candidates, []int{}, target, 0)
	return res
}
```
