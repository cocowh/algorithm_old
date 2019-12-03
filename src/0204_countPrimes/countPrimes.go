package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countPrimesByBitmap(100))
	fmt.Println(countPrimes(100))
}

func countPrimes(n int) int {
	count, m := 0, make([]bool, n)
	for i := 2; i < n; i++ {
		if !m[i] {
			count++
			for j := i + i; j < n; j += i {
				m[j] = true
			}
		}
	}
	return count
}

func countPrimesByBitmap(n int) int {
	count, size := 0, strconv.IntSize
	m := make([]int, n/size+1)
	for i := 2; i < n; i++ {
		if m[i/size]&(1<<uint(i&(size-1))) == 0 {
			count++
			for j := i + i; j < n; j += i {
				m[j/size] |= 1 << uint(j&(size-1))
			}
		}
	}
	return count
}
