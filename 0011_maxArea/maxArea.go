package _011_maxArea

func max(x, y int) int {
	if x > y {
		return  x
	} else {
		return y
	}
}

func maxArea(height []int) int {
	maxArea := 0
	i := 0
	j := len(height) - 1
	for i != j  {
		if height[i] < height[j] {
			maxArea = max(maxArea, height[i] * (j - i))
			i++
		} else {
			maxArea = max(maxArea, height[j] * (j - i))
			j--
		}
	}
	return maxArea
}