package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1,2,1,3,2,5}))
}


func singleNumber(nums []int) []int {
	m := make(map[int]bool)
	for _,v := range nums {
		if _, ok := m[v]; ok {
			delete(m,v)
		} else {
			m[v] = true
		}
	}
	ret := []int{}
	for v,_ := range m {
		ret = append(ret,v)
	}
	return ret
}