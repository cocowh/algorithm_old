package _530_getMinimumDifference

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type res struct {
	minDiff, previous int
}

func getMinimumDifference(root *TreeNode) int {
	res := res{0, 0}
	helper(root, &res)
	return res.minDiff
}

func helper(node *TreeNode, res *res) {
	if node == nil {
		return
	}
	helper(node.Left, res)
	newDiff := diff(res.previous, node.Val)
	if res.minDiff > newDiff {
		res.minDiff = newDiff
	}
	res.previous = node.Val
	helper(node.Right, res)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
