package main

import "fmt"

func main() {
	fmt.Println(singleNumberByHash([]int{2,2,2,3}))
}

func singleNumber(nums []int) int {
	one,two := 0,0
	for _,v := range nums {
		one = (one ^ v) &^ two
		two = (two ^ v) &^ one
 	}
	return one
}


func singleNumberByHash(nums []int) int {
	m,sum1,sum2 := make(map[int]bool),0,0
	for _,v := range nums {
		if _,ok := m[v]; !ok {
			sum1 += v
			m[v] = true
		}
		sum2 += v
	}
	return (3 * sum1 - sum2) / 2
}