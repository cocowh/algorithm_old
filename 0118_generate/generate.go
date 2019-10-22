package main

import "fmt"

func main() {
	fmt.Println(generate(10))
}
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
