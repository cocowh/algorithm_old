package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, str string) bool {
	lenS := len(pattern)
	strArr := strings.Split(str, " ")
	if len(strArr) != lenS {
		return false
	}
	for i := 0; i < lenS; i++ {
		if strings.IndexByte(pattern, pattern[i]) != indexOf(strArr, strArr[i]) {
			return false
		}
	}
	return true
}
func indexOf(arr []string,search string) int {
	for index,v := range arr {
		if v == search {
			return index
		}
	}
	return -1
}
