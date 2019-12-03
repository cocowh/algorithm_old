# [组合总和](https://leetcode-cn.com/problems/combination-sum/)

### 题目

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

* 所有数字（包括 target）都是正整数。  
* 解集不能包含重复的组合。 

示例 1:

>输入: candidates = [2,3,6,7], target = 7,    
所求解集为:
```
[
  [7],
  [2,2,3]
]
```

示例 2:

>输入: candidates = [2,3,5], target = 8,  
所求解集为:

```
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
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
		return
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		stack = append(stack, candidates[i])
		findCombinationSum(ret, candidates, stack, target-candidates[i], i)
		stack = stack[:len(stack)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	findCombinationSum(&ret, candidates, []int{}, target, 0)
	return ret
}
```
