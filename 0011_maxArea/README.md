# [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

### 题目

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

![图片](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

 

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49


### 解法

最容易想到暴力，枚举所有结果取最大值，n*(n-1)/2，O(n^2)

然而并非所有枚举都是有必要的，area=width*height

只需保证max(width)和max(height)

因此考虑由两边向中间枚举计算最大值，只需O(n)

```
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
```