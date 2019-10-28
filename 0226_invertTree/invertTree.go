package _226_invertTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	invertTreeByRecursive(root)
	return root
}

func invertTreeByRecursive(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
	invertTreeByRecursive(node.Left)
	invertTreeByRecursive(node.Right)
}