# [同构字符串](https://leetcode-cn.com/problems/isomorphic-strings/)

### 题目

给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:

>输入: s = "egg", t = "add"  
输出: true  

示例 2:

>输入: s = "foo", t = "bar"  
输出: false  

示例 3:

>输入: s = "paper", t = "title"  
输出: true  

说明:  
你可以假设 s 和 t 具有相同的长度。

### 解法

首先想到的是map互相映射比较，然后参考[评论区](https://leetcode-cn.com/problems/isomorphic-strings/solution/1-xing-python-by-knifezhu-5/)，还有首次索引比较这种骚操作。

```
func isIsomorphic(s string, t string) bool {
	ma, mb := make(map[byte]byte, len(s)), make(map[byte]byte, len(s))
	for index, sb := range s {
		if v1, ok := ma[byte(sb)]; ok {
			if v1 != t[index] {
				return false
			}
		} else if v2, ok := mb[t[index]]; ok {
			if v2 != byte(sb) {
				return false
			}
		} else {
			ma[s[index]] = t[index]
			mb[t[index]] = s[index]
		}
	}
	return true
}

func isIsomorphicByArrIndex(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
}

func isIsomorphicByIndex(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Index(s, string(s[i])) != strings.Index(t,string(t[i])) {
			return false
		}
	}
	return true
}
```
