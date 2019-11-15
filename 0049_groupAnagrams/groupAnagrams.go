package main

import (
	"fmt"
	"sort"
)

func main() {
	//sortStr := []byte("qweqdagwerqegert")
	//sort.Slice(sortStr, func(i, j int) bool {
	//	return sortStr[i] < sortStr[j]
	//})
	//fmt.Println(string(sortStr))
	//fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagramsBySort([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagramsBySort(strs []string) [][]string {
	m := make(map[string][]string,len(strs))
	for _,str := range strs {
		sortStr := []byte(str)
		sort.Slice(sortStr, func(i, j int) bool {
			return sortStr[i] < sortStr[j]
		})
		if _, ok := m[string(sortStr)]; ok {
			m[string(sortStr)] = append(m[string(sortStr)], str)
		} else {
			m[string(sortStr)] = []string{str}
		}
	}
	res := make([][]string,0,len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sign := make([]int, 26)
		for _, char := range str {
			sign[char-'a']++
		}
		signStr := make([]byte,0,52)
		for i := 0; i < 26; i++ {
			signStr = append(signStr, '#')
			signStr = append(signStr, byte(sign[i] + '0' - 0))
		}
		fmt.Println(string(signStr))
		if _, ok := m[string(signStr)]; ok {
			m[string(signStr)] = append(m[string(signStr)], str)
		} else {
			m[string(signStr)] = []string{str}
		}
	}
	fmt.Println(m)
	res := make([][]string,0,len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
