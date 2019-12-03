package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElementByVote([]int{2,2,1,1,1,2,2}))
	fmt.Println(majorityElementByVote([]int{2,2,1,1,1,2,2}))
}

/*
 投票
 */
func majorityElementByVote(nums []int) int {
	vote,target := 0,nums[0]
	for _,value := range nums {
		if value == target {
			vote++
		} else {
			vote --
			if vote < 0 {
				 target = value
				 vote = 0
			}
		}
	}
	return target
}

func majorityElementByMap(nums []int) int {
	m,target := make(map[int]int),nums[0]
	for _,value := range nums {
		if _,ok := m[value];ok {
			m[value]++
		} else {
			m[value] = 1
		}
		if m[value] > m[target] {
			target = value
		}
	}
	return target
}

func majorityElementBySort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

