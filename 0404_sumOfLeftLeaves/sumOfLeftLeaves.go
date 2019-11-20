package _404_sumOfLeftLeaves

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func helper(node *TreeNode, flag bool, res int) int  {
	if node == nil {
		return res
	}
	if node.Right == nil && node.Left == nil && flag {
		return res + node.Val
	}
	return helper(node.Left, true, res) + helper(node.Right, false, res)
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root.Left, true, 0) + helper(root.Right, false, 0)
}

