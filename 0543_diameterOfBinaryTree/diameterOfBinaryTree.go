package _543_diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(node *TreeNode) (length, res int) {
	if node == nil {
		return
	}

	leftLen, leftDia := helper(node.Left)
	rightLen, rightDia := helper(node.Right)

	length = max(leftLen, rightLen) + 1
	res = max(leftLen+rightLen, max(leftDia, rightDia))
	return
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}
