package _016_threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ret := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l,r := i + 1,length-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if math.Abs(float64(target - sum)) < math.Abs(float64(target - ret)) {
				ret = sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}