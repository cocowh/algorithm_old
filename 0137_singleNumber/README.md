# [只出现一次的数字 II](https://leetcode-cn.com/problems/single-number-ii/)

### 题目

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

>输入: [2,2,3,2]  
输出: 3  
示例 2:

>输入: [0,1,0,1,0,1,99]  
输出: 99

### 解法

容易想到hash

参考[异或三进位](https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0137.single-number-ii/single-number-ii.go)

```
func singleNumber(nums []int) int {
	one,two := 0,0
	for _,v := range nums {
		one = (one ^ v) &^ two
		two = (two ^ v) &^ one
 	}
	return one
}


func singleNumberByHash(nums []int) int {
	m,sum1,sum2 := make(map[int]bool),0,0
	for _,v := range nums {
		if _,ok := m[v]; !ok {
			sum1 += v
			m[v] = true
		}
		sum2 += v
	}
	return (3 * sum1 - sum2) / 2
}
```
