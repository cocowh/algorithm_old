# [跳跃游戏 II](https://leetcode-cn.com/problems/jump-game/)

### 题目

给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

```
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
```

说明:

假设你总是可以到达数组的最后一个位置。

### 解法

参考[贪心](https://leetcode-cn.com/problems/jump-game-ii/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-10/)。

```
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
```
