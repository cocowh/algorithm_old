package main

import "fmt"

func main()  {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
	ret,sum,lenN := nums[0],0,len(nums)
	for i := 0; i < lenN; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > ret {
			ret = sum
		}
	}
	return ret
}