# [杨辉三角](https://leetcode-cn.com/problems/pascals-triangle/)

### 题目

给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

>输入: 5  
输出:  
[  
     [1],  
    [1,1],  
   [1,2,1],  
  [1,3,3,1],  
 [1,4,6,4,1]  
]  

### 解法

go切片操作

```
func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}
	res = append(res,[]int{1})
	for i := 1; i < numRows; i++ {
		row := []int{}
		prevRow := res[i-1]
		row = append(row, 1)
		for j := 1; j < i; j++ {
			row = append(row, prevRow[j-1]+prevRow[j])
		}
		row = append(row, 1)
		res = append(res, row)
	}
	return res
}
```
