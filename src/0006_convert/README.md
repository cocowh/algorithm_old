# [Z字形变换](https://leetcode-cn.com/problems/zigzag-conversion/)
### 题目

将字符串 "PAYPALISHIRING" 以Z字形排列成给定的行数：

```
P   A   H   N
A P L S I I G
Y   I   R
```
之后从左往右，逐行读取字符："PAHNAPLSIIGYIR"

实现一个将字符串进行指定行数变换的函数:

string convert(string s, int numRows);
示例 1:

输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
示例 2:

输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
解释:

```
P     I    N
A   L S  I G
Y A   H R
P     I
```
### 解法
想到的遍历字符串确定位置，然后拼接，或者按确定位置在字符串遍历然后拼接。

```
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	gap := 2*numRows - 2
	var ret string
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j += gap {
			ret += string(s[i+j])
			if i != 0 && i != numRows-1 && j+gap-i < length {
				ret += string(s[j+gap-i])
			}
		}
	}
	return ret
}
```

