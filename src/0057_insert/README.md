# [插入区间](https://leetcode-cn.com/problems/insert-interval/)

### 题目

给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

```
输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
```

示例 2:

```
输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
```
### 解法

```
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
```
