package _104_maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthSimple(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthSimple(root.Left), maxDepthSimple(root.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depth++
	return max(helper(node.Left, depth), helper(node.Right, depth))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
