package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 3, 5, 2, 9}))
}

func backtrack(length, first int, ret *[][]int, nums []int) {
	if first == length {
		tmp := make([]int,length)
		copy(tmp,nums)
		*ret = append(*ret, tmp)
	}
	for i := first; i < length; i++ {
		nums[first],nums[i] = nums[i],nums[first]
		backtrack(length,first+1,ret,nums)
		nums[first],nums[i] = nums[i],nums[first]
	}
}

func permute(nums []int) [][]int {
	ret := [][]int{}
	numsTemp := nums
	backtrack(len(numsTemp), 0, &ret, numsTemp)
	return ret
}
