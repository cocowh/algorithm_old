package _003_lengthOfLongestSubstring

import "math"

//解法1 暴力破解

func lengthOfLongestSubstring1(s string) int {
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

//解法2 滑动窗口

func lengthOfLongestSubstring2(s string) int {
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

//解法3 优化的滑动窗口

func lengthOfLongestSubstring3(s string) int {
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