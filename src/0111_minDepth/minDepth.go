package _111_minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDep := int(^uint(0) >> 1)
	if root.Left != nil {
		minDep = min(minDepth(root.Left), minDep)
	}
	if root.Right != nil {
		minDep = min(minDepth(root.Right), minDep)
	}
	return minDep + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
