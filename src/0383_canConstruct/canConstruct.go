package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aba"))
}

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
