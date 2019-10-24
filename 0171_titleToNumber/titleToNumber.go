package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("Z"))
}

func titleToNumber(s string) int {
	len,ret := len(s),0
	for i := len - 1; i >=0; i-- {
		ret += int(s[i] - 'A' + 1) * int(math.Pow(26 , float64(len - 1 - i)))
	}
	return ret
}

func titleToNumberFromHead(s string) int {
	step,ret := 0,0
	for i := len(s) - 1; i >=0; i-- {
		ret += int(s[i] - 'A' + 1) * int(math.Pow(26 , float64(step)))
		step++
	}
	return ret
}