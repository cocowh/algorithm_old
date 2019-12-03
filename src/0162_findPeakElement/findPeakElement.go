package main

import "fmt"

func main() {
	//fmt.Println(findPeakElement([]int{3, 2, 1}))
	fmt.Println(findPeakElementByBinary([]int{1,2,3}))
}

func findPeakElement(nums []int) int {
	length := len(nums)
	head, tail := 1, length-2
	for head <= tail {
		if nums[head-1] < nums[head] && nums[head] > nums[head+1] {
			return head
		} else {
			head++
		}
		if nums[tail-1] < nums[tail] && nums[tail] > nums[tail+1] {
			return tail
		} else {
			tail--
		}
	}
	if nums[0] > nums[length-1] {
		return 0
	} else {
		return length - 1
	}
}

func findPeakElementByBinary(nums []int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) / 2
		if nums[mid] > nums[mid+1] {
			tail = mid
		} else {
			head = mid + 1
		}
	}
	return head
}
