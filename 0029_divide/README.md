# [两数相除](https://leetcode-cn.com/problems/divide-two-integers/)

### 题目
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

### 解法

思路一：

 * abs(dividend)-abs(divisor) > 0 quotient++
 * dividend^divisor确定符号
 * 面对边界超时不通过
     
思路二：

 * 除法由dividend高位 - divisor >= 0的商quotient高位
 * dividend高位向低位移位，例如十进制移位 999 / 5，999 >> 2 - 5 >0 ，有商高位quotient += 1 << 2 => 100, 以此类推
 * 可有若abs(dividend) >> n >= abs(divisor)，则quotient += 1 << n
 * 对32位有符号整数，最多遍历移位32次，不会超时


```
func divide(dividend int, divisor int) int {
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))
	quotient := 0
	for i := uint32(31);i >= 0 && dividendAbs >= divisorAbs; i-- {
		if (dividendAbs >> i) >= divisorAbs {
			quotient = quotient + ( 1 << i)
			dividendAbs = dividendAbs - (divisorAbs << i)
		}
	}
	if dividend ^ divisor < 0 {
		quotient = -quotient
	}
	if quotient > math.MaxInt32 || quotient < math.MinInt32 {
		return  math.MaxInt32
	}
	return quotient
}
```