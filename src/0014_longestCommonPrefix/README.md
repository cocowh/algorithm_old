# [最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

### 题目

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

### 解法

    求最长公共前缀，遍历长度以最短字符串为准。go可以对字符串进行切片，求最长前缀可以从前往后取最长字串横向比较是否相等。

    遇到不相等可取长度直接往前进一位，再次横向依次比较，遇到相等判断是否比较到最后一个字符串字串，若是则此时为最长公共前缀。

    若n个字符串最短长度m，最坏时间复杂度O(n*m)，期盼类型是有公共前缀的用例。

    看到有从前往后依次横向比较字符的解法，最坏时间复杂度也是O(n*m)，但是期盼的是无公共前缀的用例，字符比较要比字符串比较性能更好些，更推荐从前往后解法吧。


```
func longestCommonPrefix(strs []string) string {
	strArrLen := len(strs)
	if strArrLen == 0 {
		return ""
	}
	if strArrLen == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for _,str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for minLen >0 {
		for i := 0; i< strArrLen; i++ {
			if strs[i][:minLen] == strs[0][:minLen] {
				if  i == strArrLen - 1 {
					return strs[0][:minLen]
				} else {
					continue
				}
			} else {
				minLen--
				break
			}
		}
	}
	return  ""
}
```