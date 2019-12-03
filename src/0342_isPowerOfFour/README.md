# [4的幂](https://leetcode-cn.com/problems/power-of-four/)

### 题目
给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

示例 1:

>输入: 16  
输出: true  

示例 2:

>输入: 5  
输出: false  


进阶：
你能不使用循环或者递归来完成本题吗？

### 解法

递归爆栈了，然后参考(位运算解法)[https://leetcode-cn.com/problems/power-of-four/solution/wei-shi-yao-xian-pan-duan-shi-fou-wei-2de-mi-by-va/]。

```
func isPowerOfFour(num int) bool {
	if num < 0 || num & (num - 1) != 0 {
		return false
	}
	return num & 0x55555555 != 0
}
```
