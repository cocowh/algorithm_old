package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
	fmt.Println(permuteUnique([]int{0, 1, 0, 0, 9}))
}

func backtrack(res *[][]int, nums, stack []int, token []bool, length, start int) {
	if start == length {
		temp := make([]int, len(stack))
		copy(temp, stack)
		*res = append(*res, temp)
		return
	}
	used := make(map[int]bool, length-start)
	for i := 0; i < length; i++ {
		if !token[i] && !used[nums[i]] {
			used[nums[i]] = true
			token[i] = true
			stack[start] = nums[i]
			backtrack(res, nums, stack, token, length, start+1)
			token[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	numsTemp := nums
	sort.Ints(numsTemp)
	token := make([]bool, len(numsTemp))
	stack := make([]int,len(numsTemp))
	backtrack(&res, numsTemp, stack, token, len(numsTemp), 0)
	return res
}
