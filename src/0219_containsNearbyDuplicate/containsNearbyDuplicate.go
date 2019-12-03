package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for index, v := range nums {
		if i,ok := m[v];ok && index - i <= k {
			return true
		} else {
			m[v] = index
		}
	}
	return false
}
