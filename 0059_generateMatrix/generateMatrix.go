package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

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
