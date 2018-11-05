# [最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)
### 题目

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

### 解法一
参考：[中心扩展算法](https://leetcode-cn.com/problems/longest-palindromic-substring/solution/)


```
func longestPalindrome(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < length; {
		if length-i <= maxLen/2 {
			break
		}
		b, e := i, i
		for e < length-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1
		for e < length-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > maxLen {
			start = b
			maxLen = newLen
		}
	}
	return s[start : start+maxLen]
}
```

