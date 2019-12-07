# [字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

### 题目

给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

```
s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
```
注意事项：您可以假定该字符串只包含小写字母。

### 解法
同[赎金信](https://github.com/cocowh/algorithm/tree/master/src/0383_canConstruct)

```
func firstUniqChar(s string) int {
	m := map[int32]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}
	for _, v := range s {
		m[v]++
	}
	for index, v := range s {
		if m[v] == 1 {
			return index
		}
	}
	return -1
}

func firstUniqCharBetter(s string) int {
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}
	for index, v := range s {
		if m[v-'a'] == 1 {
			return index
		}
	}
	return -1
}
```
