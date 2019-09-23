package main

import "fmt"

func main() {
	fmt.Println(mySqrtNew(100))
}

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	low,high := 0,x
	for low < high {
		mid := (low + high) / 2
		if mid * mid > x {
			high = mid
		} else if mid * mid < x {
			low = mid
		} else {
			return mid
		}
		if high - low == 1 {
			return  low
		}
		fmt.Println(low,high,mid)
	}
	return low
}

func mySqrtNew(x int) int {
	res := x
	for res * res > x {
		res = (res + x / res) / 2
	}
	return res
}
