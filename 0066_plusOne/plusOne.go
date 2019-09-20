package main

import "fmt"

func main()  {
	fmt.Println(plusOne([]int{1,2,3}))
}

func plusOne(digits []int) []int {
	length := len(digits)
	c := 0
	for i := length - 1; i >= 0; i-- {
		sum := digits[i] + c + 1
		c := sum / 10
		digits[i] = sum % 10
		if c == 0 {
			return digits
		}
		if i == 0 && c > 0 {
			digits = append([]int{c}, digits...)
		}
	}
	return digits
}