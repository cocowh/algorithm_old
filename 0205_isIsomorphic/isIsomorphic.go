package main

import (
	"fmt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/strings"
)

func main() {
	test := [][]string{
		[]string{"egg", "add"},
		[]string{"foo", "bar"},
		[]string{"paper", "title"},
	}
	for _, v := range test {
		fmt.Println(isIsomorphicByIndex(v[0], v[1]))
	}
	//fmt.Println(isIsomorphic("ab","aa"))
}

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
