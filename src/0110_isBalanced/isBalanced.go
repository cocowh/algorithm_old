package _110_isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	flag, _ := helper(root)
	return flag
}

func helper(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	flagL, levelL := helper(node.Left)
	flagR, levelR := helper(node.Right)
	if flagL && flagR && levelL-levelR >= -1 && levelL-levelR <= 1 {
		return true, max(levelR, levelL) + 1
	}
	return false, 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
