package main

import (
	"fmt"
)

func main() {
	data := []int{1,0,0,1}
	//moveZeroes(data)
	moveZeroesByCount(data)
	fmt.Println(data)
}

func moveZeroesByCount(nums []int) {
	count := 0
	for i,value := range nums {
		if value == 0 {
			count++
		} else {
			nums[ i - count] = nums[i]
		}
	}
	for i := len(nums) - count ; i < len(nums) ; i++ {
		nums[i] = 0
	}
}

func moveZeroes(nums []int)  {
	zeroPo := 0
	for i := 0; i < len(nums) ; i++ {
		if nums[i] != 0 {
			if i != zeroPo {
				nums[zeroPo] = nums[i]
				nums[i] = 0
			}
			zeroPo++
		}
	}
}
