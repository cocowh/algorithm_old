package _055_canJump

func canJump(nums []int) bool {
	length := len(nums)
	lastPos := length - 1
	for i := length - 1; i >= 0 ; i-- {
		if i + nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
