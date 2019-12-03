# [字符串相加](https://leetcode-cn.com/problems/add-strings/)

### 题目

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

1. num1 和num2 的长度都小于 5100.
2. num1 和num2 都只包含数字 0-9.
3. num1 和num2 都不包含任何前导零。
4. 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

### 解法
时空效率不佳。

[更好的方式](https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0415.add-strings/add-strings.go)

```
func addStrings(num1 string, num2 string) string {
	res, gap, i, j, carry := []byte{}, int('0'-0), len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 || carry > 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i]) - gap
		}
		if j >= 0 {
			n2 = int(num2[j]) - gap
		}
		temp := n1 + n2 + carry
		carry = temp / 10
		res = append([]byte{byte(temp % 10 + gap)},res...)
		i--
		j--
	}
	return string(res)
}
```
