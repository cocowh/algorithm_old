package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	h, t := 0, len(numbers)-1
	for h < t {
		sum := numbers[h] + numbers[t]
		if sum == target {
			return []int{h + 1, t + 1}
		} else if sum < target {
			h++
		} else {
			t--
		}
	}
	return []int{}
}
