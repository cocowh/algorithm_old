package main

import "fmt"

func main() {
	intervals := [][]int{
		[]int{1, 5},
	}
	newInterval := []int{2, 3}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 || len(newInterval) == 0 {
		return [][]int{newInterval}
	}
	//首
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	//尾
	if newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}
	res := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		if (intervals[i][0] <= newInterval[0] && newInterval[0] <= intervals[i][1]) || (intervals[i][0] <= newInterval[1] && newInterval[1] <= intervals[i][1]) || (newInterval[0] <= intervals[i][0] && intervals[i][0] <= newInterval[1]) {
			if intervals[i][0] < newInterval[0] {
				newInterval[0] = intervals[i][0]
			}
			if intervals[i][1] > newInterval[1] {
				newInterval[1] = intervals[i][1]
			}
			if i == len(intervals)-1 {
				res = append(res, newInterval)
			}
		}
		if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
		}
		if intervals[i][0] >= newInterval[1] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			break
		}
	}
	return res
}
