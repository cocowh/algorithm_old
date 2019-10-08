package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(' ', byte('a'), byte('z'), 'A','Z','0','9')
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)
	if length <= 1 {
		 return true
	}
	left,right := 0,length - 1
	for left < right {
		for !judgeEffectiveChar(s[left]) && left < right {
			left++
		}
		for !judgeEffectiveChar(s[right]) && right > left{
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return  true
}

func judgeEffectiveChar(b byte) bool {
	if  ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}