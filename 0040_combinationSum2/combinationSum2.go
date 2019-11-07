package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
}

func findCombinationSum(ret *[][]int, candidates, stack []int, target, start int) {
	if target == 0 {
		temp := make([]int, len(stack))
		copy(temp, stack)
		*ret = append(*ret, temp)
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		stack = append(stack, candidates[i])
		findCombinationSum(ret, candidates, stack, target-candidates[i], i+1)
		stack = stack[:len(stack)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	findCombinationSum(&res, candidates, []int{}, target, 0)
	return res
}
