package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(15))
}

func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res += n / 5
		n /= 5
	}
	return res
}
