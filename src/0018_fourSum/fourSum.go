package _018_fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length - 3; i++{
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < length - 2; j++ {
			if j > i +1 && nums[j] == nums[j-1] {
				continue
			}
			l,r := j + 1,length - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target || (l != j + 1 &&  nums[l] == nums[l - 1]){
					l++
				} else if sum > target  || (r != length - 1 && nums[r] == nums[r + 1]){
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return ret
}