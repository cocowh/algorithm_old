package _034_searchRange

func searchRange(nums []int, target int) []int {
	start,end,length := -1,-1,len(nums)
	if length == 0 || nums[0] > target || nums[length - 1] < target{
		return  []int{start, end}
	} else if length == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{start, end}
		}
	}
	left,right := 0,length - 1
	for left <= right && ( end == -1 || start == -1) {
		if (start != -1 && end != -1) {
			break
		}
		if nums[left] == target  && start == -1{
			start = left
		}
		if nums[right] == target && end == -1 {
			end = right
		}
		mid := (left + right) / 2
		if nums[mid] > target && end == -1 {
			right = mid - 1
		}
		if nums[mid] < target && start == -1 {
			left = mid + 1
		}
		if nums[mid] == target {
			if start == -1 {
				left = left + 1
			}
			if end == -1 {
				right = right - 1
			}
		}
	}
	return  []int{start, end}
}