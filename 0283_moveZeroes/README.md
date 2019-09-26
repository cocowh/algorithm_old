# [移动零](https://leetcode-cn.com/problems/move-zeroes/)

### 题目

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

>输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

### 解法

随简单比较有意思

最简单思路：冒泡，遇0遍历后移。双层遍历时间复杂度O(n^2)。

快慢指针：慢遇0停，快继续，快慢指针不指向同一个位置且快不为0时，交换快慢指针值，慢指针前进。

累计移位补0：累计遍历时遇0次数，遇0位置之前的相对位置替换为当前位置值，后续遇0长度补0。
   
```
func moveZeroesByCount(nums []int) {
	count := 0
	for i,value := range nums {
		if value == 0 {
			count++
		} else {
			nums[ i - count] = nums[i]
		}
	}
	for i := len(nums) - count ; i < len(nums) ; i++ {
		nums[i] = 0
	}
}

func moveZeroes(nums []int)  {
	zeroPo := 0
	for i := 0; i < len(nums) ; i++ {
		if nums[i] != 0 {
			if i != zeroPo {
				nums[zeroPo] = nums[i]
				nums[i] = 0
			}
			zeroPo++
		}
	}
}
```