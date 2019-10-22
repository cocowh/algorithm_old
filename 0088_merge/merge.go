package main

import "fmt"

func main()  {
	nums1,m,nums2,n :=  []int{1,2,3,0,0,0},3,[]int{2,5,6},3
	merge(nums1,m,nums2,n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 {
		copy(nums1,nums2)
		return
	}
	p,p1,p2 := m + n - 1, m - 1, n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	if p2 >= 0 {
		copy(nums1[:p+1],nums2[:p2+1])
	}
}