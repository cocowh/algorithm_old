package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 1}))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if m[v] {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}

func containsDuplicateBySort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
