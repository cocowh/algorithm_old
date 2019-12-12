# [检测大写字母](https://leetcode-cn.com/problems/detect-capital/)
### 题目

给定一个单词，你需要判断单词的大写使用是否正确。

我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如"USA"。
单词中所有字母都不是大写，比如"leetcode"。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
否则，我们定义这个单词没有正确使用大写字母。

示例 1:

```
输入: "USA"
输出: True
```

示例 2:

```
输入: "FlaG"
输出: False  
```

### 解法

比较好的思路[大写字母的个数小于正在遍历的下标](https://leetcode-cn.com/problems/detect-capital/solution/c-ji-jian-jie-fa-by-makeex/)

```
func detectCapitalUse(word string) bool {
	count, lenW := 0, len(word)
	for index, v := range word {
		if v <= 'Z' && v >= 'A' {
			if count < index {
				return false
			}
			count++
		}
	}
	if count == lenW || count <= 1 {
		return true
	}
	return false
}
```
