# [螺旋矩阵](https://leetcode-cn.com/problems/spiral-matrix/)

### 题目

给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

```
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
```
示例 2:

```
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
```
### 解法

上下左右按层遍历

```
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	minRow, maxRow, minCol, maxCol, res := 0, row-1, 0, col-1, []int{}
	for minRow <= maxRow && minCol <= maxCol {
        //上层
		for c := minCol; c <= maxCol; c++ {
			res = append(res, matrix[minRow][c])
		}
        //右层
		for r := minRow + 1; r <= maxRow; r++ {
			res = append(res, matrix[r][maxCol])
		}
		if minRow < maxRow && minCol < maxCol {
            //下层
			for c := maxCol - 1; c > minCol; c-- {
				res = append(res, matrix[maxRow][c])
			}
            //左层
			for r := maxRow; r > minRow; r-- {
				res = append(res, matrix[r][minCol])
			}
		}
        //收缩
		minCol++
		minRow++
		maxCol--
		maxRow--
	}
	return res
}
```
