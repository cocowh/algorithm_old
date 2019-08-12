package _015_threeSum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	length := len(nums)
	if length == 0 || nums[0] > 0 || nums[length - 1] < 0 {
		return ret
	}
	for i := 0; i < length; i++ {
		if nums[i] > 0{
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l,r := i + 1,length-1
		for l < r {
			if nums[l] + nums[r] == -nums[i]{
				if l == i + 1 || r == length - 1 {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
					l++
					r--
				} else if nums[l] == nums[l - 1] {
					l++
				} else if nums[r] == nums[r + 1] {
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
					l++
					r--
				}
			} else if nums[l] + nums[r] < -nums[i] {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}