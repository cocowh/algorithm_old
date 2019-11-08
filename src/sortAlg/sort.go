package main

import (
	"fmt"
)

func main() {
	nums := []int{9,4,1,5,3,8}
	//fastSort([]int{1, 0, 2, 9, 5})
	sortByBigHeap(nums)
	fmt.Println(nums)
}

func SortMain() {
	var n int
	fmt.Println("输入数组大小（n）：")
	fmt.Scanf("%d", &n)
	var arr []int
	fmt.Printf("请依次输入%d个数（中间空格或换行）\n", n)
	for index := 0; index < n; index++ {
		var temp int
		fmt.Scanf("%d", &temp)
		arr = append(arr, temp)
	}
	fmt.Printf("输入的数组为%v\n", arr)
	fmt.Println("****插入排序****")
	InsertionSort(arr)
	fmt.Println("****插入排序****")
	fmt.Printf("输入的数组为%v\n", arr)
	fmt.Println("****归并排序****")
	MergeSort(arr)
	fmt.Println("****归并排序****")
	fmt.Printf("输入的数组为%v\n", arr)
	fmt.Println("****冒泡排序****")
	BubbleSort(arr)
	fmt.Println("****冒泡排序****")
	fastSort(arr)
}

//插入排序
func InsertionSort(arr []int) {
	len := len(arr)
	newArr := make([]int, len)
	copy(newArr, arr)
	for i := 1; i < len; i++ {
		temp := newArr[i]
		j := i
		for ; j > 0 && newArr[j-1] > temp; j-- {
			newArr[j] = newArr[j-1]
		}
		newArr[j] = temp
		fmt.Printf("第%d次插入排序结果为：%v\n", i, newArr)
	}
}

//归并排序
func MergeSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	num := len / 2
	left := MergeSort(arr[:num])
	right := MergeSort(arr[num:])
	return merge(left, right)
}

//合并
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	fmt.Printf("每次的归并输出堆：%v\n", result)
	return
}

//冒泡排序
func BubbleSort(arr []int) {
	len := len(arr)
	newArr := make([]int, len)
	copy(newArr, arr)
	for i := 1; i < len; i++ {
		flag := false
		for j := 1; j < len; j++ {
			var key = newArr[j]
			if newArr[j-1] > newArr[j] {
				newArr[j] = newArr[j-1]
				newArr[j-1] = key
				flag = true
			}
		}
		if !flag {
			break
		}
		fmt.Printf("第%d次冒泡排序结果为：%v\n", i, newArr)
	}
}

//快排

func fastSort(arr []int) {
	left, right := 0, len(arr)-1
	fastHelper(arr, left, right)
	fmt.Println(arr)
}

func fastHelper(arr []int, left, right int) {
	p, temp, i, j := left, arr[left], left, right
	for i <= j {
		for j >= p && arr[j] >= temp {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		for i <= p && arr[i] <= temp {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
		arr[p] = temp
		if p-left > 1 {
			fastHelper(arr, left, p-1)
		}
		if right-p > 1 {
			fastHelper(arr, p+1, right)
		}
	}
}

//最大堆排序

func sortByBigHeap(nums []int) {
	maxHeapify(nums)
	fmt.Println(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}
}

func maxHeapify(nums []int) {
	for i := 0; i < len(nums); i++ {
		curIndex := i
		fatIndex := (curIndex - 1) / 2
		for nums[curIndex] > nums[fatIndex] {
			nums[curIndex], nums[fatIndex] = nums[fatIndex], nums[curIndex]
			curIndex = fatIndex
			fatIndex = (curIndex - 1) / 2
		}
	}
}

func heapify(nums []int, index, size int) {
	left, right := 2*index+1, 2*index+2
	for left < size {
		maxNumIndex := 0
		if nums[left] < nums[right] && right < size {
			maxNumIndex = right
		} else {
			maxNumIndex = left
		}
		if nums[index] > nums[maxNumIndex] {
			maxNumIndex = index
		}
		if index == maxNumIndex {
			break
		}
		nums[maxNumIndex], nums[index] = nums[index], nums[maxNumIndex]
		index = 2*index + 1
		right = 2*index + 2
	}
}
