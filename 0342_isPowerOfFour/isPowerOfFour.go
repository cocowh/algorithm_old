package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfFour(0))
}

func isPowerOfFour(num int) bool {
	if num < 0 || num & (num - 1) != 0 {
		return false
	}
	return num & 0x55555555 != 0
}