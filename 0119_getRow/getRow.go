package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRowBetter(4))
}

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