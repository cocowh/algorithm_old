# [整数拆分](https://leetcode-cn.com/problems/integer-break/)

### 题目
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

>输入: 2  
输出: 1

解释: 2 = 1 + 1, 1 × 1 = 1。

示例 2:

>输入: 10  
输出: 36

解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

### 解法

[参考](https://leetcode-cn.com/problems/integer-break/solution/)

```
func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3

	if b == 0 {
		return int(math.Pow(float64(3), float64(a)))
	}
	switch b {
	case 0:
		return int(math.Pow(float64(3), float64(a)))
	case 1:
		return int(math.Pow(float64(3), float64(a-1))) * 4
	default:
		return int(math.Pow(float64(3), float64(a))) * 2
	}
}

func integerBreakByHash(n int) int {
	m := map[int]int{
		2:1,
	}
	for i := 3; i<= n; i++ {
		for j := 1; j <= i -1; j++ {
			m[i] = int(math.Max(float64(m[i]), math.Max(float64( j * m[i - j]), float64(j * (i - j)))))
		}
	}
	return m[n]
}
```
