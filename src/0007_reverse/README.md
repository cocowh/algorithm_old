# [反转整数](https://leetcode-cn.com/problems/reverse-integer/)
### 题目
给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。

### 解法
思考：根据题意，整数值范围在[-214748368,2147483647]，大于math.MaxInt32或者小于math.MinInt32为溢出，返回0。对原值依次除10取余，商为下次迭代初始值，结果值依次乘10加上所得余数，直至原值除尽为0。在期间判断值是否超出范围。


```
func reverse(x int) int {
	ret := 0
	for x != 0 {
		rem := x % 10 //求余
		x /= 10
		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && rem > 7) {
			return 0
		}
		if ret < math.MinInt32/10 || (ret == math.MinInt32/10 && rem < -8) {
			return 0
		}
		ret = ret*10 + rem
	}
	return ret
}
```

