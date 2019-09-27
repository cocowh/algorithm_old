package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
}

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1;i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}