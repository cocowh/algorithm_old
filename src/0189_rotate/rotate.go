package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	//rotate(nums, 2)
	rotateByReverse(nums,2)
	fmt.Println(nums)
}

//错误解法
func rotateWrong(nums []int, k int) {
	length := len(nums)
	if length <= 1 || k == 0 {
		return
	}
	if length%k != 0 || k == 1 {
		count, point, temp := 0, 0, nums[0]
		for count < length {
			point = (point + k) % length
			cur := nums[point]
			nums[point] = temp
			temp = cur
			count++
		}
	} else {
		for i := 0; i < k; i++ {
			loopW := length / k
			temp := nums[(loopW*k+i)%length]
			for loopW > 0 {
				fmt.Println("lo", loopW)
				cur := nums[(loopW-1)*k+i]
				nums[(loopW-1)*k+i] = temp
				temp = cur
				loopW--
			}
		}
	}
}


func rotate(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		cur, pre, cycle := start, nums[start], false
		for start != cur || !cycle {
			cycle = true
			next := (cur + k) % len(nums)
			temp := nums[next]
			nums[next] = pre
			pre = temp
			cur = next
			count++
		}
	}
}

func rotateByReverse(nums []int, k int) {
	k = k % len(nums)
	reverseArr(nums, 0, len(nums)-1)
	reverseArr(nums, 0, k-1)
	reverseArr(nums, k, len(nums)-1)
}

func reverseArr(nums []int, start, end int) {
	for start < end {
		nums[start],nums[end] = nums[end],nums[start]
		start++
		end--
	}
}
