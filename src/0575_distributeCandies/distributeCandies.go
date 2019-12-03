package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distributeCandies([]int{1,2,3,1,4,5}))
}

func distributeCandies(candies []int) int {
	//sister := make(map[int]bool)优化，指定初始大小，避免动态分配
	sister := make(map[int]bool,len(candies))
	for _,v := range candies {
		if _, ok := sister[v];ok {
			continue
		} else {
			sister[v] = true
		}
	}
	if len(sister) > len(candies)/2 {
		return  len(candies) / 2
	} else {
		return  len(sister)
	}
}

//效率不高
func distributeCandiesAnother(candies []int) int {
	length := len(candies)
	if length == 0 {
		return 0
	}
	sort.Ints(candies)
	count := 1
	for i := 1; i < length ; i++ {
		if candies[i] != candies[i-1] {
			count++
		}
		if count > length / 2  {
			return length / 2
		}
	}
	return count
}