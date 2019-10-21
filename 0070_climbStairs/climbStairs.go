package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	ret := make(map[int]int)
	ret[1] = 1
	ret[2] = 2
	for i := 3;i <= n;i++ {
		ret[i] = ret[i - 1] + ret[i - 2]
	}
	return ret[n]
}