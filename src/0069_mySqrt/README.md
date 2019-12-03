# [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)

### 题目
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
     由于返回类型是整数，小数部分将被舍去。

### 解法

简单难度。

二分。

数学解法，牛顿积分。
   
```
func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	low,high := 0,x
	for low < high {
		mid := (low + high) / 2
		if mid * mid > x {
			high = mid
		} else if mid * mid < x {
			low = mid
		} else {
			return mid
		}
		if high - low == 1 {
			return  low
		}
		fmt.Println(low,high,mid)
	}
	return low
}
```