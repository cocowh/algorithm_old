package _026_removeDuplicates

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return len(nums)
	}
	temp := nums[0]
	index := 1
	for i := 1; i < length ; i++ {
		if nums[i] == temp {
			continue
		} else {
			nums[index] = nums[i]
			temp = nums[i]
			index++
		}
	}
	return index
}
