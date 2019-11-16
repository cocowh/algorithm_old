package _098_isValidBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, -1<<63, 1<<63-1)
}
func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	return min < node.Val && node.Val < max && helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
