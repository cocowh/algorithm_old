# [串联所有单词的子串](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)

### 题目
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

 

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

### 解法

参考: [解法](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-6/)

```
func findSubstring(s string, words []string) []int {
	ret,wordNum := []int{},len(words)
	if wordNum == 0 {
		return ret
	}
	wordLen := len(words[0])
	allWards := make(map[string]int)
	for _,word := range words {
		if _,ok := allWards[word];ok {
			allWards[word]++
		} else {
			allWards[word] = 0
		}
	}
	for j := 0;j < wordLen; j++ {
		num := 0
		hasWords := make(map[string]int)
		for i := j; i < len(s) - wordNum * wordLen + 1;i = i + wordLen {
			hasRemoved := false
			for num < wordNum {
				word := s[i + num * wordLen : i + (num + 1) * wordLen]
				if _, ok := allWards[word];ok{
					if _, ok := hasWords[word]; ok {
						hasWords[word]++
					} else {
						hasWords[word] = 0
					}
					if hasWords[word] > allWards[word] {
						hasRemoved = true
						removeNum := 0
						for hasWords[word] > allWards[word] {
							firstWord := s[i + removeNum * wordLen : i + (removeNum + 1) * wordLen]
							hasWords[firstWord]--
							removeNum++
						}
						num = num - removeNum + 1
						i = i + (removeNum - 1) * wordLen
						break
					}
				} else {
					hasWords = make(map[string]int)
					i = i + num * wordLen
					num = 0
					break
				}
				num++
			}
			if num == wordNum {
				ret = append(ret,i)
			}
			if num > 0 && !hasRemoved {
				firstWord := s[i : i + wordLen]
				hasWords[firstWord]--
				num--
			}
		}
	}
	return ret
}
```