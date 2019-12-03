# [下一个排列](https://leetcode-cn.com/problems/next-permutation/)

### 题目
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

### 解法

参考: [解法](https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode/)
    
```
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
```