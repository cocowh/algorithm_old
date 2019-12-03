package _045_jump

func jump(nums []int) int {
	end, maxPosition, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i-- {
		if nums[i]+i > maxPosition {
			maxPosition = nums[i] + i
		}
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func jumpFromEnd(nums []int) int {
	position, steps := len(nums)-1, 0
	for position != 0 {
		for i := 0; i < position; i++ {
			if nums[i] + i >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}
