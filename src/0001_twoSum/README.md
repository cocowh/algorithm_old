# [两数之和](https://leetcode-cn.com/problems/two-sum/)
### 题目

给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

### 解法一
双重遍历

```
func twoSum(nums []int, target int) []int {
    length := len(nums)
	for k,v := range nums {
		for i := k + 1 ; i < length; i++ {
			if  v + nums[i] == target {
                return []int{k,i}
			}
		}
	}
   	return nil
}
```

### 解法二

存map，数组键作为map的值，数组值作为map的键，判断a+b=target即可通过map中是否存在b=target-a的键值对进行处理，此时需要一层遍历，但消耗了额外的存储空间。

```
func twoSum(nums []int, target int) []int {

	index := make(map[int]int, len(nums))

	for k,v := range nums {
		if i, ok := index[target-v]; ok {
			return []int{k,i}
		}
		index[v] = k
	}
	return  nil
}
```
