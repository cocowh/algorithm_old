package _144_preorderTraversal

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	ret := append([]int{},root.Val)
	ret = append(ret,preorderTraversal(root.Left)...)
	ret = append(ret,preorderTraversal(root.Right)...)
	return  ret
}

func preorderTraversalByStack(root *TreeNode) []int {
	ret,stack,point := []int{},[]*TreeNode{},root
	for point != nil || len(stack) > 0{
		if point != nil {
			ret = append(ret,point.Val)
			stack = append(stack,point)
			point = point.Left
		} else {
			point = stack[len(stack) - 1]
			point = point.Right
			stack = stack[:len(stack) - 1]
		}
	}
	return ret
}