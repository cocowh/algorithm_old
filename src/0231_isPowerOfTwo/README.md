# [2的幂](https://leetcode-cn.com/problems/power-of-two/)

### 题目

给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

>输入: 1  
输出: true  
解释: 20 = 1  

示例 2:

>输入: 16  
输出: true  
解释: 24 = 16  

示例 3:  

>输入: 218  
输出: false  

### 解法

做过4的幂，这题是4幂基础。判断n异或n-1是否为零。

还有一种思路，有高位到低位求与2^n余，若为0则为2的幂。循环次数有限。
   
```
func isPowerOfTwo(n int) bool {
	return n > 0 && n & (n-1) == 0
}
```