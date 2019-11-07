package main

import "fmt"

func main() {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	rotate(data)
	fmt.Println(data)
}

func rotate(matrix [][]int) {
	lenM := len(matrix)
	//i := 0; i < (lenM+1)/2 奇数时中间值旋转
	for i := 0; i < (lenM+1)/2; i++ {
		for j := 0; j < lenM/2; j++ {
			matrix[lenM-1-j][i], matrix[lenM-1-i][lenM-j-1], matrix[j][lenM-1-i], matrix[i][j] = matrix[lenM-1-i][lenM-j-1], matrix[j][lenM-1-i], matrix[i][j], matrix[lenM-1-j][i]
		}
	}
}
