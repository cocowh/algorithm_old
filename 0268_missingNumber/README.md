# [缺失数字](https://leetcode-cn.com/problems/missing-number/)

### 题目

给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:

>输入: [3,0,1]
输出: 2

示例 2:

>输入: [9,6,4,2,3,5,7,0,1]
输出: 8

说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

### 解法

首先想到高斯算法，秒过。然后看解析还有奇妙的异或求法。

参考[缺失数字](https://leetcode-cn.com/problems/missing-number/solution/que-shi-shu-zi-by-leetcode/)。
   
```
func missingNumber(nums []int) int {
	sum := 0
	for _,value := range nums {
		sum += value
	}
	return (len(nums) * (len(nums) + 1)) / 2 - sum
}

func missingNumberByXOR(nums []int) int {
	ret := len(nums)
	for index,value := range nums {
		ret ^= index ^ value
		fmt.Println(ret)
	}
	return ret
}
```