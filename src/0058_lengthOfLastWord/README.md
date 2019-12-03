# [最后一个单词的长度](https://leetcode-cn.com/problems/length-of-last-word/)

### 题目
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:

输入: "Hello World"
输出: 5

### 解法

简单难度。
   
```
func lengthOfLastWord(s string) int {
	length := len(s)
	count,flag := 0,false
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if flag {
				return count
			} else {
				continue
			}
		} else if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
			count ++
			flag = true
		} else {
			count = 0
			flag =false
		}
	}
	if flag {
		return  count
	} else {
		return 0
	}
}
```