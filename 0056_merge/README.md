# [合并区间](https://leetcode-cn.com/problems/merge-intervals/)

### 题目

给出一个区间的集合，请合并所有重叠的区间。

示例 1:

```
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
```

示例 2:

```
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
```

### 解法

对切片按第一位数排序，然后依次比较合并

golang知识点：sort.Slice()对切片排序。

```
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
```
