package main

import "fmt"

func main() {
	fmt.Println(myPow(1.0000,-2147483648))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x,n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := fastPow(x, n / 2)
	if n % 2 == 0 {
		return half * half
	}
	return half * half * x
} 