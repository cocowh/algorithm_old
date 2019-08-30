//package _036_isValidSudoku
package main
import "fmt"

func main()  {
	in := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	fmt.Println(isValidSudokuBetter(in))
}

func isValidSudoku(board [][]byte) bool {
	mr := make(map[int]map[byte]int)
	mc := make(map[int]map[byte]int)
	mb := make(map[int]map[byte]int)
	subMc := map[byte]int{}
	subMb := map[byte]int{}
	for row,cols := range board {
		subMr := map[byte]int{}
		for col,val := range cols {
			if val == '.' {
				continue
			} else {
				blockIndex := (row / 3 ) * 3 + col / 3
				_,okr := mr[row][val]
				_,okc := mc[col][val]
				_,okb := mb[blockIndex][val]
				fmt.Println("---------")
				fmt.Println(row, col, blockIndex, val, okr, okc, okb )
				fmt.Println(mr)
				fmt.Println(mc)
				fmt.Println(mb)
				if okr || okc || okb {
					return false
				} else {
					subMr[val] = 1
					mr[row] = subMr
					if _, ok := mc[col]; ok {
						subMc = mc[col]
					} else {
						subMc = map[byte]int{}
					}
					subMc[val] = 1
					mc[col] = subMc
					if _, ok := mb[blockIndex]; ok {
						subMb = mb[blockIndex]
					} else {
						subMb = map[byte]int{}
					}
					subMb[val] = 1
					mb[blockIndex] = subMb
				}
			}
		}
	}
	return true
}

func isValidSudokuBetter(board [][]byte) bool {
	mr := [9][10]bool{}
	mc := [9][10]bool{}
	mb := [9][10]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j <= 9; j++ {
			mr[i][j] = false
			mc[i][j] = false
			mb[i][j] = false
		}
	}
	for row,vals := range board {
		for col,val := range vals {
			if val == '.' {
				continue
			} else {
				val = val - 48
				blockIndex := (row / 3) * 3 + col / 3
				if mr[row][val] || mc[col][val] || mb[blockIndex][val] {
					return false
				} else {
					mr[row][val] = true
					mc[col][val] = true
					mb[blockIndex][val] = true
				}
			}
		}
	}
	return true
}