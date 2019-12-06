# [找不同](https://leetcode-cn.com/problems/find-the-difference/)

### 题目

给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。

示例:

```
输入：
s = "abcd"
t = "abcde"

输出：
e

解释：
'e' 是那个被添加的字母。
```

### 解法

累加相减+异或

哈希额外存储不考虑。

```
func findTheDifference(s string, t string) byte {
	ret := 0
	for _,v := range s {
		ret ^= int(v)
	}
	for _,v := range t {
		ret ^= int(v)
	}
	return byte(ret)
}

func findTheDifferenceByAdd(s string, t string) byte {
	counts,countt := int32(0),int32(0)
	for _,v := range s {
		counts += v
	}
	for _,v := range t {
		countt += v
	}
	return byte(countt-counts)
}
```
