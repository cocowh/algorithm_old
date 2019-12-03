package main

import (
	"fmt"
	"sort"
)

func main() {
	data := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	fmt.Println(merge(data))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) <= 1 {
		return intervals
	}
	res := [][]int{}
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			if intervals[i][1] >= temp[1] {
				temp[1] = intervals[i][1]
			}
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
		fmt.Println(res)
	}
	res = append(res, temp)
	return res
}

func mergeFistWrongSolution(intervals [][]int) [][]int {
	m, nums := make(map[int]int, len(intervals)), make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		m[intervals[i][0]] = intervals[i][1]
		nums[i] = intervals[i][0]
	}
	res := [][]int{}
	sort.Ints(nums)
	fmt.Println(nums, m)
	for i := 0; i < len(nums)-1; {
		if m[nums[i]] > nums[i+1] {
			m[nums[i]] = m[nums[i+1]]
			delete(m, nums[i+1])
			temp := nums[:i]
			nums = append(temp, nums[i+1:]...)
		} else {
			i++
		}
	}
	for i, v := range m {
		res = append(res, []int{i, v})
	}
	return res
}
