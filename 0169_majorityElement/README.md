# [求众数](https://leetcode-cn.com/problems/majority-element/)

### 题目
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2

### 解法

首先想出的遍历一遍使用hash统计所有数出现的个数。

然后看解析还有机智的排序，利用条件"众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素"。

更机智的投票算法，同样利用条件"众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素"。
   
```
/*
 投票
 */
func majorityElementByVote(nums []int) int {
	vote,target := 0,nums[0]
	for _,value := range nums {
		if value == target {
			vote++
		} else {
			vote --
			if vote < 0 {
				 target = value
				 vote = 0
			}
		}
	}
	return target
}

func majorityElementByMap(nums []int) int {
	m,target := make(map[int]int),nums[0]
	for _,value := range nums {
		if _,ok := m[value];ok {
			m[value]++
		} else {
			m[value] = 1
		}
		if m[value] > m[target] {
			target = value
		}
	}
	return target
}

func majorityElementBySort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
```