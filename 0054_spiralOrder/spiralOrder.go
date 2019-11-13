package main

import "fmt"

func main() {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(spiralOrder(data))
}

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	minRow, maxRow, minCol, maxCol, res := 0, row-1, 0, col-1, []int{}
	for minRow <= maxRow && minCol <= maxCol {
		for c := minCol; c <= maxCol; c++ {
			res = append(res, matrix[minRow][c])
		}
		for r := minRow + 1; r <= maxRow; r++ {
			res = append(res, matrix[r][maxCol])
		}
		if minRow < maxRow && minCol < maxCol {
			for c := maxCol - 1; c > minCol; c-- {
				res = append(res, matrix[maxRow][c])
			}
			for r := maxRow; r > minRow; r-- {
				res = append(res, matrix[r][minCol])
			}
		}
		minCol++
		minRow++
		maxCol--
		maxRow--
	}
	return res
}
