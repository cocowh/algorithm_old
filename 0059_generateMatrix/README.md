# [螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/)

### 题目

给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

```
输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
```
### 解法

上下左右按层遍历，同[螺旋矩阵](https://github.com/cocowh/algorithm/tree/master/0054_spiralOrder)

```
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	minRow, maxRow, minCol, maxCol, num := 0, n-1, 0, n-1, 0
	for minRow <= maxRow && minCol <= maxCol {
		for c := minCol; c <= maxCol; c++ {
			num++
			res[minRow][c] = num
		}
		for r := minRow + 1; r <= maxRow; r++ {
			num++
			res[r][maxCol] = num
		}
		if minRow < maxRow && minCol < maxCol {
			for c := maxCol - 1; c > minCol; c-- {
				num++
				res[maxRow][c] = num
			}
			for r := maxRow; r > minRow; r-- {
				num++
				res[r][minCol] = num
			}
		}
		minCol++
		minRow++
		maxCol--
		maxRow--
	}
	return res
}
```
