package _007_reverse

import "math"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		rem := x % 10 //求余
		x /= 10
		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && rem > 7) {
			return 0
		}
		if ret < math.MinInt32/10 || (ret == math.MinInt32/10 && rem < -8) {
			return 0
		}
		ret = ret*10 + rem
	}
	return ret
}
