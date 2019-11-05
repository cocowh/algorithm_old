package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
	fmt.Println(numTreesByFormula(3))
}

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func numTreesByFormula(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = ret * 2 * (2*i + 1) / (i + 2)
	}
	return ret
}
