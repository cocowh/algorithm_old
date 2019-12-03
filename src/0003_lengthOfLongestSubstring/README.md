# [无重复字符的最长字串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
### 题目

给定一个字符串，找出不含有重复字符的最长子串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 无重复字符的最长子串是 "abc"，其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。

### 解法
容易想到的是暴力破解

参考[官方解体](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/)的滑动窗口及优化的滑动窗口

#### 暴力破解

思路：双层遍历原字符串取子字符串，对子字符串进行是否为无重复字符串判断，是则取其长度，和原长度比较取最大值。


```
func lengthOfLongestSubstring(s string) int {
	len := len(s)
	maxLen := 0
	for i := 0; i < len;i++{
		for j := i + 1; j <= len; j++{
			if allunique(s, i, j) {
				math.Max(float64(maxLen), float64(j-i))
			}
		}
	}
	return maxLen
}

func allunique(s string, start int, end int) bool {
	checkMap := make(map[byte]int, end-start)

	for i := start; i < end; i++ {
		if _, ok := checkMap[s[i]]; ok{
			return false
		}
		checkMap[s[i]] = i
	}
	return true
}
```

#### 滑动窗口

思路：设置集合为滑动窗口存储字符串连续不重复的字符，每次向其中添加字符时计算一次长度，与原最大长度比较取最大值。若集合中不存在字符串中的当前遍历字符，则加入到集合中，字符串下标（右标：用于遍历字符串下标）下移一位（+1），计算当前长度（右标-左标），以当前窗口左标、右标（经过+1，相当于遍历字符串对字符串的下一个字符进行比较）开始下一轮判断比较；若集合存在字符串中的当前遍历字符，则移动左标（+1），删除集合中的字符串左标对应字符，以当前左标（经过+1，相当于删除左标至右标之间的元素直至集合中不出现当前正被比较的右标字符）、右标开始下一轮比较；直至左标或者右标达到字符串长度表明窗口移动到末尾。

```
func lengthOfLongestSubstring(s string) int {
	len := len(s)
	positionMap := make(map[byte]int, len)
	maxLen := 0

	for j,i := 0, 0;j < len && i < len; {
		if _,ok := positionMap[s[j]]; !ok {
			j++
			positionMap[s[j]] = j
			maxLen = int(math.Max(float64(maxLen), float64(j - i)))
		} else {
			i++
			delete(positionMap, s[i])
		}
	}
	return maxLen
}
```

#### 优化的滑动窗口

思路：上一个算法对一个字符的比较操作可能经过很多次的移除添加操作，最终的目的是获取当前字符与其上一次出现位置之间的子字符串。所以可以使用map存储字符，键为字符，值为其出现的位置。若当前遍历的字符（窗口右标）在map中出现，取其上一次出现的位置（确定窗口左标）。每次计算长度（右标-坐标+1），与原长度比较取最大值。

```
func lengthOfLongestSubstring(s string) int {
	len := len(s)
	positionMap := make(map[byte]int, len)

	maxLen := 0

	for j, i := 0, 0;j < len;j++ {
		if _,ok := positionMap[s[i]]; ok {
			i = int(math.Max(float64(positionMap[s[j]]), float64(i)))
		}
		maxLen = int(math.Max(float64(maxLen), float64(j-i+1)))
		positionMap[s[j]] = j+1
	}
	return maxLen
}
```