package main

import (
	"fmt"
	"sort"
)

func main() {
	data, target := []int{
		1, 2, 3, 4, 5,
	}, 4
	fmt.Println(combinationSum(data, target))
}

func findCombinationSum(ret *[][]int, candidates, stack []int, target, start int) {
	if target == 0 {
		temp := make([]int, len(stack))
		copy(temp, stack)
		*ret = append(*ret, temp)
		return
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		stack = append(stack, candidates[i])
		findCombinationSum(ret, candidates, stack, target-candidates[i], i)
		stack = stack[:len(stack)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	findCombinationSum(&ret, candidates, []int{}, target, 0)
	return ret
}
