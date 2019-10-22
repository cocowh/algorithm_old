# [杨辉三角](https://leetcode-cn.com/problems/pascals-triangle/)

### 题目

给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

![图片](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

>输入: 3
输出: [1,3,3,1]

进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

### 解法

go切片操作

解法一：参考[杨辉三角](https://github.com/cocowh/algorithm/tree/master/0118_generate)，根据上一行求下一行

解法二：数学公式求解，[参考](https://leetcode-cn.com/problems/pascals-triangle-ii/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--28/)

```
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	lRow:= []int{1}
	for i := 1; i <= rowIndex; i++ {
		curRow := []int{}
		curRow = append(curRow, 1)
		for j := 1; j < i; j++ {
			curRow = append(curRow, lRow[j-1]+lRow[j])
		}
		curRow = append(curRow, 1)
		lRow = curRow
	}
	return lRow
}

func getRowBetter(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	ans,pre := []int{},1
	ans = append(ans,1)
	for i := 1; i <= rowIndex; i++ {
		cur := pre * (rowIndex - i + 1)/i
		ans = append(ans,cur)
		pre = cur
	}
	return ans
}
```
