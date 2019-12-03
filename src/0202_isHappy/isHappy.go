package main

import "fmt"

func main() {
	fmt.Println(powSum(powSum(2)),powSum(2),isHappy(10))
}

func isHappy(n int) bool {
	fast,slow := powSum(powSum(n)),powSum(n)
	for fast != slow {
		fmt.Println(fast,slow)
		if fast == 1 || slow == 1{
			return true
		}
		fast,slow = powSum(powSum(fast)),powSum(slow)
	}
	return slow == 1
}

func powSum(n int) int {
	ret := 0
	for n > 0 {
		ret += (n % 10) * (n % 10)
		n = n / 10
	}
	return ret
}
