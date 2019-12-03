package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1,2,3,1}))
}

func rob(nums []int) int {
	prevMax, currMax := 0, 0
	for _, v := range nums {
		if prevMax + v  > currMax{
			currMax,prevMax = prevMax + v,currMax
		}  else {
			prevMax = currMax
		}
	}
	return currMax
}
