package _027_removeElement

func removeElement(nums []int, val int) int {
	index := 0
	length := len(nums)
	for i := 0; i < length; i++{
		if nums[i] == val {
			continue
		} else {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func removeElementTWO(nums []int, val int) int {
	for i := 0; i < len(nums); i++{
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}