package _343_integerBreak

import "math"

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3

	if b == 0 {
		return int(math.Pow(float64(3), float64(a)))
	}
	switch b {
	case 0:
		return int(math.Pow(float64(3), float64(a)))
	case 1:
		return int(math.Pow(float64(3), float64(a-1))) * 4
	default:
		return int(math.Pow(float64(3), float64(a))) * 2
	}
}

func integerBreakByHash(n int) int {
	m := map[int]int{
		2:1,
	}
	for i := 3; i<= n; i++ {
		for j := 1; j <= i -1; j++ {
			m[i] = int(math.Max(float64(m[i]), math.Max(float64( j * m[i - j]), float64(j * (i - j)))))
		}
	}
	return m[n]
}