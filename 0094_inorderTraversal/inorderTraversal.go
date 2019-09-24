package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func inorderTraversalByStack(root *TreeNode) []int {
	res,stack,point := []int{},[]*TreeNode{},root
	for point != nil || len(stack) > 0{
		for point != nil {
			stack = append(stack,point)
			point = point.Left
		}
		point = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res,point.Val)
		point = point.Right
	}
	return res
}