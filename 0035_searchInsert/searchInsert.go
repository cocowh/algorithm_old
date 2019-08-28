package _035_searchInsert

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length  == 0 || nums[0] > target{
		return 0
	} else if nums [length - 1]  < target {
		return length
	}
	left,right := 0,length - 1
	for left <= right {
		if nums[left] == nums[right] {
			if nums[left] >= target {
				return left
			} else {
				return left + 1
			}
		}
		if nums[left] ==target {
			return left
		}
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] == target {
			left++
		}
	}
	if nums[left] >= target {
		return left
	} else {
		return left + 1
	}
}
