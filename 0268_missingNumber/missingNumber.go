package main

import "fmt"

func main()  {
	//fmt.Println(7^3)
	fmt.Println(missingNumberByXOR([]int{3,0,1}))
}

func missingNumber(nums []int) int {
	sum := 0
	for _,value := range nums {
		sum += value
	}
	return (len(nums) * (len(nums) + 1)) / 2 - sum
}

func missingNumberByXOR(nums []int) int {
	ret := len(nums)
	for index,value := range nums {
		ret ^= index ^ value
		fmt.Println(ret)
	}
	return ret
}