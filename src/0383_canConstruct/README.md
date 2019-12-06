# [赎金信](https://leetcode-cn.com/problems/ransom-note/)

### 题目

给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。

(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)

注意：

你可以假设两个字符串均只含有小写字母。

```
canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
```

### 解法

```
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int, len(ransomNote)/2)
	for _, v := range magazine {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range ransomNote {
		if num, ok := m[v]; ok && num > 0 {
			m[v]--
		} else {
			return false
		}
	}
	return true
}

//优化
func canConstructBetter(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		m[v-'a']--
		if m[v-'a'] < 0 {
			return false
		}
	}
	return true
}
```
