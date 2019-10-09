# [Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)

### 题目

实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

>输入: 2.00000, 10  
输出: 1024.00000  

示例 2:  

>输入: 2.10000, 3  
输出: 9.26100  

示例 3:

>输入: 2.00000, -2  
输出: 0.25000  
解释: 2-2 = 1/22 = 1/4 = 0.25  

说明:  
-100.0 < x < 100.0  
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

### 解法

一开自然想到暴力循环，然后超时没通过。

然后就是想到幂等拆解了，但是怎么拆解没有思路，最后参考了二分拆解[快速幂等](https://leetcode-cn.com/problems/powx-n/solution/powx-n-by-leetcode/)。

```
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x,n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := fastPow(x, n / 2)
	if n % 2 == 0 {
		return half * half
	}
	return half * half * x
} 
```