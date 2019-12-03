package main

import "fmt"

func main()  {
	fmt.Println(findDuplicate([]int{1,2,3,3}))
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	ret := nums[0]
	for ret != slow {
		ret = nums[ret]
		slow = nums[slow]
	}
	return ret
}