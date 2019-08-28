package _033_search

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0  {
		return -1
	}
	if length == 1{
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	pointIndex := 0
	if nums[0] > nums[length - 1] {
		for left,right := 0,length - 1;left <= right; {
			mid := (left + right) / 2
			if nums[mid] > nums[mid + 1] {
				pointIndex = mid + 1
				break
			} else if nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if nums[pointIndex] == target {
		return pointIndex
	} else if pointIndex == 0 {
		return binarySearch(nums,target,0,length - 1)
	} else if nums[0] > target {
		return binarySearch(nums, target,pointIndex + 1, length -1)
	} else {
		return binarySearch(nums, target, 0, pointIndex - 1)
	}
}

func binarySearch(nums []int, target,left,right int) int {
	for  left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}