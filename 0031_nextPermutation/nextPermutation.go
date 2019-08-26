package _031_nextPermutation

func nextPermutation(nums []int)  {
	length := len(nums)
	if length <= 1 {
		return
	}
	left := length - 1
	for ; left > 0 ; left-- {
		if nums[left] <= nums[left - 1]  {
			continue
		} else {
			break
		}
	}
	if left == 1 && nums[left] <= nums[left - 1]  || left == 0 {
		for i := 0; i < length / 2; i++ {
			temp := nums[i]
			nums[i] = nums[length - 1 - i]
			nums[length - 1 -i] = temp
		}
	} else {
		left--
		for right := length - 1; right > left; right-- {
			if nums[right] > nums[left] {
				temp := nums[right]
				nums[right]  = nums[left]
				nums[left] = temp
				break
			}
		}
		i,j := left + 1,length - 1
		for i < j {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
			j--
		}
	}
}